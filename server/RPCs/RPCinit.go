package RPCs

import (
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	pb "raftAlgo.com/service/server/gRPC"
)

type logEntry struct {
	command string
	term    int
}

//var term , serverID , lastLogIndex , lastLogTerm, commitIndex, lastApplied, leaderId int64 = 2,2,2,2,2,2,2
const (
    leader = 2
    candidate = 1
    follower = 0
)

var State = follower

type server struct {
	pb.UnimplementedRPCServiceServer
	currentTerm  int64
	serverId     int64
	lastLogIndex int64
	lastLogTerm  int64
	commitIndex  int64
	lastApplied  int64
	candidateId  int64
	votedFor     int64
	nextIndex    []int64
	matchIndex   []int64
	log          []*pb.RequestAppendLogEntry
	leaderId     int64
}

func (s *server) initServerDS() {
	serverID, _ := strconv.Atoi(os.Getenv("CandidateID"))
	serverId := int64(serverID)
	s.currentTerm = 0
	s.serverId = serverId
	s.lastLogIndex = 1 // Update in Future
	s.lastLogTerm = 1
	s.commitIndex = 0
	s.lastApplied = 0
	s.candidateId = 0
	s.votedFor = 0
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	s.nextIndex = make([]int64, REPLICAS)
	s.matchIndex = make([]int64, REPLICAS)
	s.leaderId = 0
}

func (s *server) initLeaderDS() (bool){
    response := false
    candidateId :=  os.Getenv("CandidateID")
	CandidateID, _ := strconv.Atoi(candidateId)
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	log.Printf("Server %v : initLeaderDS : Setting Candidate State to Leader State ", candidateId)
	if State==candidate {
        State = leader
        logLength := len(s.log)
        for i:=0 ; i<REPLICAS; i++ {
            s.nextIndex[i] = int64(logLength)
        }
        s.leaderId = int64(CandidateID)
        response = true
    } else {
        log.Printf("Server %v : initLeaderDS : Unable to set Candidate State to Leader State : Current State : %v ", candidateId,State)
        response = false
    }
    return response
}

func (s *server) initCandidateDS() (bool) {
    response := false
    candidateId :=  os.Getenv("CandidateID")
    log.Printf("Server %v : initCandidateDS : Setting follower State to Candidate State ", candidateId)
    if State==follower{
        State = candidate
        response = true
    } else {
        log.Printf("Server %v : initCandidateDS : Unable to set follower State to Leader State : Current State : %v ", candidateId,State)
        response = false
    }
    return response
}

func (s *server) initFollowerDS() (bool) {
    response := true
    candidateId :=  os.Getenv("CandidateID")
    oldState := State
    State = follower // No matter what was your state get Back to follower
    log.Printf("Server %v : initFollowerDS : Setting Current State : %v to follower State : %v", candidateId,oldState,State)
    return response
}


func RPCInit() bool {
	port := ":" + os.Getenv("PORT") + os.Getenv("CandidateID")
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : Port ID for the current Server : %v", serverId, port)
	serverobj := server{}
	serverobj.initServerDS()
	go serverobj.ElectionInit()
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
