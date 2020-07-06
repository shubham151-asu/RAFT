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

<<<<<<< HEAD
func (s *server) RequestVoteRPC(ctx context.Context, in *pb.RequestVote) (*pb.ResponseVote, error) {
	serverId :=  os.Getenv("CandidateID")
	term := in.GetTerm()
	candidateId := in.GetCandidateID()
	lastLogTerm := in.GetLastLogTerm()
	lastLogIndex := in.GetLastLogIndex()
	log.Printf("Server %v : Received Term : %v",serverId, term)
	log.Printf("Server %v : Received CandidateID : %v", serverId, candidateId)
	log.Printf("Server %v : Received LastLogIndex : %v", serverId, lastLogIndex)
	log.Printf("Server %v : Received LastLogTerm : %v", serverId, lastLogTerm)
	//TODO Add code to response for RequestVote

	switch {
	    case s.currentTerm>term:
	        return &pb.ResponseVote{Term:s.currentTerm,VoteGranted:false}, nil
	    case (!s.votedFor  || candidateId>=0) && (s.lastLogTerm>= lastLogTerm && s.lastLogIndex>=lastLogIndex): // This condition needs to be verified
	        log.Printf("Server %v : vote granted to : %v",serverId, candidateId) // Do Additional things
	        return &pb.ResponseVote{Term:s.currentTerm,VoteGranted:true}, nil
	    }
	return &pb.ResponseVote{Term:s.currentTerm,VoteGranted:false}, nil
}

func (s *server) VoteRPC(address string) (bool){
	serverId :=  os.Getenv("CandidateID")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Server %v : could not connect : error %v", serverId, err)
	}
	defer conn.Close()
	c := pb.NewRPCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.RequestVoteRPC(ctx, &pb.RequestVote{Term: s.currentTerm,CandidateID:s.serverId,
	                                                LastLogIndex:s.lastLogIndex,LastLogTerm:s.lastLogTerm})
	if err != nil {
		log.Fatalf("Server %v : could not Receive Vote : error %v", serverId, err)
	}
	log.Printf("Server %v : Response received %s",serverId, response.String())

	// TODO Update server currentTerm in all responses

	return response.GetVoteGranted()
=======



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
>>>>>>> 6757736b6d636bd5c46348b259e3d12daa33dfc7
}
