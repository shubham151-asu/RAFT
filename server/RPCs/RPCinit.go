package RPCs

import (
	"database/sql"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"

	pb "raftAlgo.com/service/server/gRPC"
)

type logEntry struct {
	logIndex int
	command  string
	term     int
}

//var term , serverID , lastLogIndex , lastLogTerm, commitIndex, lastApplied, leaderId int64 = 2,2,2,2,2,2,2
const (
	leader    = 2
	candidate = 1
	follower  = 0
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
	leaderId int64
	Lock     sync.Mutex
	db       *sql.DB
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

func (s *server) getLastLog() (index, term int64) {
	s.Lock.Lock()
	index = s.lastLogIndex
	term = s.lastLogTerm
	s.Lock.Unlock()
	return index, term
}

func (s *server) setLastLog(index, term int64) {
	s.Lock.Lock()
	s.lastLogIndex = index
	s.lastLogTerm = term
	s.Lock.Unlock()
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

func RPCInit() bool {
	port := ":" + os.Getenv("PORT") + os.Getenv("CandidateID")
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : Port ID for the current Server : %v", serverId, port)
	serverobj := server{}
	serverobj.initServerDS()
	go serverobj.ElectionInit()
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
	db, err := sql.Open("sqlite3", "./logger.db") //TODO rename db
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	s.db = db
	s.createDBStructure()
}

func (s *server) createDBStructure() {
	log.Printf("createDBStructure")
	statement, err := s.db.Prepare("CREATE TABLE IF NOT EXISTS logs (logIndex INTEGER PRIMARY KEY, term INTEGER, command TEXT)")
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement.Exec()
	statement, err = s.db.Prepare("CREATE TABLE IF NOT EXISTS common (property TEXT PRIMARY KEY, value INTEGER)")
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement.Exec()

	rows, err := s.db.Query("SELECT property, value FROM common")
	if err != nil {
		log.Fatal("err : ", err)
	}
	if !rows.Next() {
		log.Printf("No data in common")
		statement, err = s.db.Prepare("INSERT INTO common (property, value) VALUES (?, ?)")
		if err != nil {
			log.Fatal("err : ", err)
		}
		statement.Exec("currentTerm", 0)
		statement.Exec("votedFor", 0)
	}
	rows, err = s.db.Query("SELECT logIndex,term FROM logs WHERE logIndex = (SELECT MAX(logIndex) FROM logs)")
	if err != nil {
		log.Fatal("err : ", err)
	}
	if rows.Next() {
		log.Printf("log exist")
		var logIndex int
		var term int
		rows.Scan(&logIndex, &term)
		s.lastLogIndex = int64(logIndex)
		s.lastLogTerm = int64(term)
		log.Printf("lastLogIndex : " + strconv.Itoa(logIndex) + "     lastLogTerm:" + strconv.Itoa(term))
	} else {
		log.Printf("No data in logs")
	}
}
func (s *server) insertLog(logIndex1 int, term1 int, command1 string) {
	log.Printf("insertLog --> logIndex:" + strconv.Itoa(int(logIndex1)) + " term:" + strconv.Itoa(int(term1)) + " Command: " + command1)
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement, err := s.db.Prepare("INSERT INTO logs (logIndex,term, command) VALUES (?, ?,?)")
	if err != nil {
		log.Fatal("err : ", err)
	}
	_, err = tx.Stmt(statement).Exec(logIndex1, term1, command1)
	if err != nil {
		log.Printf("doing rollback")
		tx.Rollback()
	} else {
		tx.Commit()
		log.Printf("transaction commited")
	}

	statement, err = s.db.Prepare("SELECT term,command FROM logs WHERE logIndex = ?")
	rows, err := statement.Query(logIndex1)
	if err != nil {
		log.Fatal("err in getLog : ", err)
	}
	var term int
	var command string
	if rows.Next() {
		rows.Scan(&term, &command)
		log.Printf("lastLogTerm:" + strconv.Itoa(term) + "  command: " + command)
	}

	// LastInsertId, err := res.LastInsertId()
	// RowsAffected, err := res.RowsAffected()
	// log.Printf("LastInsertId : %v  RowsAffected: %v", LastInsertId, RowsAffected)
}
func (s *server) getLogList(startLogIndex int, endLogIndex int) []*pb.RequestAppendLogEntry { //inclusive
	log.Printf("startLogIndex:" + strconv.Itoa(startLogIndex) + " endLogIndex:" + strconv.Itoa(endLogIndex))
	statement, err := s.db.Prepare("SELECT logIndex,term,command FROM logs WHERE logs.logIndex >= ? AND logs.logIndex <= ?")
	rows, err := statement.Query(startLogIndex, endLogIndex)
	var response []*pb.RequestAppendLogEntry
	if err != nil {
		log.Printf("err in getLogList : ", err)
		return response
	}
	if rows.Next() {
		var logIndex int
		var term int
		var command string
		rows.Scan(&logIndex, &term, &command)
		log.Printf("lastLogIndex : " + strconv.Itoa(logIndex) + "     lastLogTerm:" + strconv.Itoa(term) + "  command: " + command)
		response = append(response, &pb.RequestAppendLogEntry{Command: command, Term: int64(term)})
	} else {
		log.Printf("No data")
	}
	return response
}

func (s *server) getLog(logIndex int) *pb.RequestAppendLogEntry {
	statement, err := s.db.Prepare("SELECT term,command FROM logs WHERE logIndex = ?")
	rows, err := statement.Query(logIndex)
	if err != nil {
		log.Fatal("err in getLog : ", err)
	}
	var term int
	var command string
	if rows.Next() {
		rows.Scan(&term, &command)
		log.Printf("lastLogTerm:" + strconv.Itoa(term) + "  command: " + command)
	}
	return &pb.RequestAppendLogEntry{Command: command, Term: int64(term)}

}
