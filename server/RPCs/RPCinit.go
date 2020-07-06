package RPCs

import
(
<<<<<<< HEAD
	"os"
	"log"
	"net"
	"strconv"
	"google.golang.org/grpc"
	pb "raftAlgo.com/service/server/gRPC"
=======
	pb "raftAlgo.com/service/server/gRPC"
	// "raftAlgo.com/service/server/RPCs"
	"os"
	"log"
	"net"
	"google.golang.org/grpc"
>>>>>>> 6757736b6d636bd5c46348b259e3d12daa33dfc7
)


type logEntry struct {
    command string
        term  int
}

<<<<<<< HEAD
//var term , serverID , lastLogIndex , lastLogTerm, commitIndex, lastApplied, leaderId int64 = 2,2,2,2,2,2,2

var nextIndex , matchIndex []int64

type server struct {
    pb.UnimplementedRPCServiceServer
    currentTerm int64
    serverId int64
    lastLogIndex int64
    lastLogTerm int64
    commitIndex int64
    lastApplied int64
    candidateId int64
    votedFor bool // Done for convinecne
}

func (s *server)initServerDS() {
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
=======
var term , serverID , lastLogIndex , lastLogTerm, commitIndex, lastApplied, leaderId int64 = 2,2,2,2,2,2,2

var nextIndex , matchIndex []int64



type server struct {
        pb.UnimplementedRPCServiceServer
>>>>>>> 6757736b6d636bd5c46348b259e3d12daa33dfc7
}

func RPCInit() bool {
	port := ":"+os.Getenv("PORT")+os.Getenv("CandidateID")
<<<<<<< HEAD
	serverId :=  os.Getenv("CandidateID")
    log.Printf("Server %v : Port ID for the current Server : %v",serverId,port)
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
=======
        log.Printf("Port ID for the current candidate : %v",port)
        lis, err := net.Listen("tcp", port)
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pb.RegisterRPCServiceServer(s, &server{})
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
>>>>>>> 6757736b6d636bd5c46348b259e3d12daa33dfc7
	return true
}
