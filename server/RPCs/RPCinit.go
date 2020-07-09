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

type server struct {
	pb.UnimplementedRPCServiceServer
	currentTerm  int64
	serverId     int64
	lastLogIndex int64
	lastLogTerm  int64
	commitIndex  int64
	lastApplied  int64
	candidateId  int64
	votedFor     bool // Done for convinecne
	nextIndex    []int64
	matchIndex   []int64
	log          []*pb.RequestAppendLogEntry
	leaderId     int64
}

func (s *server) initServerDS() {
	serverID, _ := strconv.Atoi(os.Getenv("CandidateID"))
	serverId := int64(serverID)
	s.currentTerm = 1
	s.serverId = serverId
	s.lastLogIndex = 1 // Update in Future
	s.lastLogTerm = 1
	s.commitIndex = 0
	s.lastApplied = 0
	s.candidateId = 0
	s.votedFor = false
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	s.nextIndex = make([]int64, REPLICAS)
	s.matchIndex = make([]int64, REPLICAS)
	s.leaderId = 1
}

func RPCInit() bool {
	port := ":" + os.Getenv("PORT") + os.Getenv("CandidateID")
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : Port ID for the current Server : %v", serverId, port)
	serverobj := server{}
	serverobj.initServerDS()
	//go serverobj.ElectionInit()
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
