package RPCs

import
(
	pb "raftAlgo.com/service/server/gRPC"
	// "raftAlgo.com/service/server/RPCs"
	"os"
	"log"
	"net"
	"google.golang.org/grpc"
)


type logEntry struct {
    command string
        term  int
}

var term , serverID , lastLogIndex , lastLogTerm, commitIndex, lastApplied, leaderId int64 = 2,2,2,2,2,2,2

var nextIndex , matchIndex []int64



type server struct {
        pb.UnimplementedRPCServiceServer
}

func RPCInit() bool {
	port := ":"+os.Getenv("PORT")+os.Getenv("CandidateID")
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
	return true
}
