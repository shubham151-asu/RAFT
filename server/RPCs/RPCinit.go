package RPCs

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"raftAlgo.com/service/server/DB"
	pb "raftAlgo.com/service/server/gRPC"
)

type logEntry struct {
	logIndex int
	command  string
	term     int
}

//var term , serverID , lastLogIndex , lastLogTerm, commitIndex, lastApplied, leaderId int64 = 2,2,2,2,2,2,2
const (
	leader                = 2
	candidate             = 1
	follower              = 0
	followerWriteWaitTime = 5000
)

type server struct {
	pb.UnimplementedRPCServiceServer
	currentTerm  int64
	serverId     int64
	state        int32
	lastLogIndex int64
	lastLogTerm  int64
	commitIndex  int64
	lastApplied  int64
	candidateId  int64
	votedFor     int64
	nextIndex    []int64
	matchIndex   []int64
	//log          []*pb.RequestAppendLogEntry
	leaderId     int64
	Lock         sync.Mutex
	db           DB.Conn
	stateMachine map[string]string
}

func (s *server) getState() int32 {
	return atomic.LoadInt32(&s.state)
}

func (s *server) setState(state int) {
	atomic.StoreInt32(&s.state, int32(state))
}

func (s *server) getCurrentTerm() int64 {
	return atomic.LoadInt64(&s.currentTerm)
}

func (s *server) setCurrentTerm(term int64) {
	atomic.StoreInt64(&s.currentTerm, term)
}

func (s *server) setTermAndVotedFor(term, votedFor int64) {
	s.Lock.Lock()
	s.db.SetTermAndVotedFor(int(term), int(votedFor))
	s.currentTerm = term
	s.Lock.Unlock()
}

func (s *server) getLastLog() (index, term int64) {
	//s.Lock.Lock()
	index = s.lastLogIndex
	term = s.lastLogTerm
	//s.Lock.Unlock()
	return index, term
}

func (s *server) setLastLog(index, term int64) {
	//s.Lock.Lock()
	s.lastLogIndex = index
	s.lastLogTerm = term
	//s.Lock.Unlock()
}
func (s *server) incrementLastLogIndex(index int64) {
	atomic.AddInt64(&s.lastLogIndex, index)
}
func (s *server) verifyLastLogTermIndex(index, term int64) bool {
	response := false
	s.Lock.Lock()
	if s.lastLogTerm >= term && s.lastLogIndex >= index {
		response = true
	}
	s.Lock.Unlock()
	return response
}

func (s *server) getCommitIndex() int64 {
	return atomic.LoadInt64(&s.commitIndex)
}

func (s *server) setCommitIndex(index int64) {
	atomic.StoreInt64(&s.commitIndex, index)
}

func (s *server) IncrementCommitIndex(index int64) {
	atomic.AddInt64(&s.commitIndex, index)
}

func (s *server) getLastApplied() int64 {
	return atomic.LoadInt64(&s.lastApplied)
}

func (s *server) setLastApplied(index int64) {
	atomic.StoreInt64(&s.lastApplied, index)
}

func (s *server) initServerDS() {
	serverID, _ := strconv.Atoi(os.Getenv("CandidateID"))
	serverId := int64(serverID)
	s.currentTerm = 0
	s.serverId = serverId
	s.lastLogIndex = -1 // Update in Future
	s.lastLogTerm = 0
	s.commitIndex = 0
	s.lastApplied = 0
	s.candidateId = 0
	s.votedFor = 0
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	s.nextIndex = make([]int64, REPLICAS)
	s.matchIndex = make([]int64, REPLICAS)
	s.leaderId = 0
	s.state = follower
	s.stateMachine = make(map[string]string)
}

func (s *server) initLeaderDS() bool {
	response := false
	candidateId := os.Getenv("CandidateID")
	CandidateID, _ := strconv.Atoi(candidateId)
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	log.Printf("Server %v : initLeaderDS : Setting Candidate State to Leader State ", candidateId)
	if s.getState() == candidate {
		s.setState(leader)
		//logLength := len(s.log) // Need to apply lock
		for i := 0; i < REPLICAS; i++ {
			s.nextIndex[i] = s.lastLogIndex + 1
		}
		s.leaderId = int64(CandidateID)
		response = true
	} else {
		log.Printf("Server %v : initLeaderDS : Unable to set Candidate State to Leader State : Current State : %v ", candidateId, s.getState())
		response = false
	}
	return response
}

func (s *server) initCandidateDS() bool {
	response := false
	candidateId := os.Getenv("CandidateID")
	log.Printf("Server %v : initCandidateDS : Setting follower State to Candidate State ", candidateId)
	if s.getState() == follower {
		s.setState(candidate)
		response = true
	} else {
		log.Printf("Server %v : initCandidateDS : Unable to set follower State to Leader State : Current State : %v ", candidateId, s.getState())
		response = false
	}
	return response
}

func (s *server) initFollowerDS() bool {
	response := true
	candidateId := os.Getenv("CandidateID")
	oldState := s.getState()
	s.setState(follower) // No matter what was your state get Back to follower
	log.Printf("Server %v : initFollowerDS : Setting Current State : %v to follower State : %v", candidateId, oldState, s.getState())
	return response
}

func (s *server) AddtoStateMachine() {
	serverId := os.Getenv("CandidateID")
	for {
		if s.lastApplied < s.getCommitIndex() && (s.getState() == follower || s.getState() == leader) {
			log.Printf("Server %v : AddtoStateMachine : Applying data to State Machine ", serverId)
			entryList := s.db.GetLogList(int(s.lastApplied), int(s.getCommitIndex()))
			for i := 0; i < len(entryList); i++ {
				command := entryList[i]
				if strings.EqualFold(command.GetCommand(), "put") {
					s.stateMachine[command.GetKey()] = command.GetValue()
				}
				s.lastApplied++
			}
		}
		log.Printf("Server %v : AddtoStateMachine : Nothing to add, Going to Sleep ", serverId)
		time.Sleep(time.Duration(followerWriteWaitTime) * time.Millisecond)
	}
}

func RPCInit() bool {
	port := ":" + os.Getenv("PORT") + os.Getenv("CandidateID")
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : Port ID for the current Server : %v", serverId, port)
	serverobj := server{}
	serverobj.initServerDS()
	go serverobj.ElectionInit()
	go serverobj.AddtoStateMachine()
	serverobj.DBInit()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRPCServiceServer(s, &serverobj)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return true
}

func (s *server) DBInit() {
	s.db.DBInit()
	lastLogIndex, lastLogTerm := s.db.CreateDBStructure()
	s.lastLogTerm = lastLogTerm
	s.lastLogIndex = lastLogIndex
}
