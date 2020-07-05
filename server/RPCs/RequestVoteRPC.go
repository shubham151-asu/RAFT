package RPCs

import
(
	"log"
	"time"
	"os"
	"google.golang.org/grpc"
	pb "raftAlgo.com/service/server/gRPC"
	"context"
)




func (s *server) RequestVoteRPC(ctx context.Context, in *pb.RequestVote) (*pb.ResponseVote, error) {
	serverId :=  os.Getenv("CandidateID")
	log.Printf("Server %v : Received Term : %v",serverId, in.GetTerm())
	log.Printf("Server %v : Received CandidateID : %v", serverId, in.GetCandidateID())
	log.Printf("Server %v : Received LastLogIndex : %v", serverId, in.GetLastLogIndex())
	log.Printf("Server %v : Received LastLogTerm : %v", serverId, in.GetLastLogTerm())
	//TODO Add code to response for RequestVote
	return &pb.ResponseVote{Term:term,VoteGranted:false}, nil
}

func (s *server) VoteRPC(address string) {
	serverId :=  os.Getenv("CandidateID")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRPCServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// lastLogIndex = len(logs)
	// if lastLogIndex > 0 {
	// 	lastLogTerm = logs[lastLogIndex-1].term + 1
	// } else {
	// 	lastLogTerm = 1
	// }
	response, err := c.RequestVoteRPC(ctx, &pb.RequestVote{Term: term,CandidateID:serverID,
	                                                LastLogIndex:lastLogIndex,LastLogTerm:lastLogTerm})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Server %v : Response recieved %s",serverId, response.String())
}
