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
	term := in.GetTerm()
	candidateId := in.GetCandidateID()
	lastLogTerm := in.GetLastLogTerm()
	lastLogIndex := in.GetLastLogIndex()
	log.Printf("Server %v : RequestVoteRPC : Received Term : %v",serverId, term)
	log.Printf("Server %v : RequestVoteRPC : CandidateID : %v", serverId, candidateId)
	log.Printf("Server %v : RequestVoteRPC : LastLogIndex : %v", serverId, lastLogIndex)
	log.Printf("Server %v : RequestVoteRPC : LastLogTerm : %v", serverId, lastLogTerm)
	//TODO Add code to response for RequestVote

	switch {
	    case s.getCurrentTerm()>term:
	        return &pb.ResponseVote{Term:s.currentTerm,VoteGranted:false}, nil
	    case s.verifyLastLogTermIndex(lastLogIndex,lastLogTerm):
	        if s.votedFor==0 {  // Need to verify from persistance DB
	            s.votedFor = candidateId  // Need to add to persistance
	            log.Printf("Server %v : RequestVoteRPC :vote granted to %v for term %v",serverId, candidateId,s.currentTerm) // Do Additional things
	            return &pb.ResponseVote{Term:s.getCurrentTerm(),VoteGranted:true}, nil
	        } else {
	            switch {
	               case s.getCurrentTerm()==term:
	                   log.Printf("Server %v : RequestVoteRPC : vote not granted to candidate %v as already voted to %v for the current term : %v",serverId, candidateId,s.votedFor,s.currentTerm) // Do Additional things
	                   return &pb.ResponseVote{Term:s.getCurrentTerm(),VoteGranted:false}, nil
	               case term>s.getCurrentTerm():
	                   s.votedFor = candidateId
	                   s.setCurrentTerm(term)
	                   log.Printf("Server %v : RequestVoteRPC : vote granted to %v for term %v",serverId, candidateId,s.currentTerm) // Do Additional things
	                   return &pb.ResponseVote{Term:s.getCurrentTerm(),VoteGranted:true}, nil
	            }
	        }
    }
    return &pb.ResponseVote{Term:s.getCurrentTerm(),VoteGranted:false}, nil
}

func (s *server) VoteRPC(address string) (bool){
    response := false
    serverId :=  os.Getenv("CandidateID")
    log.Printf("Server %v : VoteRPC : Current State : %v", serverId, s.getState())
    if s.getState()==candidate {
        conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil {
             log.Printf("Server %v : VoteRPC : could not connect : error %v", serverId, err)
        }
        defer conn.Close()
        c := pb.NewRPCServiceClient(conn)
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        response, err := c.RequestVoteRPC(ctx, &pb.RequestVote{Term: s.currentTerm,CandidateID:s.serverId,
                                                   LastLogIndex:s.lastLogIndex,LastLogTerm:s.lastLogTerm})
        if err != nil {
           log.Printf("Server %v : VoteRPC : could not Receive Vote : error %v", serverId, err)
        }

        log.Printf("Server %v : VoteRPC : Response received %s",serverId, response.String())
        return response.GetVoteGranted()
        // TODO Update server currentTerm in all responses
    } else {
        log.Printf("Server %v : VoteRPC : No Longer a Candidate State : Current State : %v",serverId,s.getState())
    }
	return response
}
