package RPCs

import (
	"context"
	"log"
	"math"
	"os"
	"strconv"
	"time"
	"google.golang.org/grpc"
	pb "raftAlgo.com/service/server/gRPC"
)

func (s *server) RequestAppendRPC(ctx context.Context, in *pb.RequestAppend) (*pb.ResponseAppend, error) {
	serverID := os.Getenv("CandidateID")
	log.Printf("Server %v : RequestAppendRPC : Received term : %v", serverID, in.GetTerm())
	log.Printf("Server %v : RequestAppendRPC : Received leaderId : %v", serverID, in.GetLeaderId())
	log.Printf("Server %v : RequestAppendRPC : Received prevLogIndex : %v", serverID, in.GetPrevLogIndex())
	log.Printf("Server %v : RequestAppendRPC : Received prevLogTerm : %v", serverID, in.GetPrevLogTerm())
	// TODO reset timer
	s.ResetTimer()
	if in.GetTerm() < s.currentTerm {
		return &pb.ResponseAppend{Term: s.currentTerm, Success: false}, nil
	}
	log.Printf("Server %v : RequestAppendRPC : Length of log : %v", serverID , len(s.log))
	if in.GetPrevLogIndex() >= 0 && (int64(len(s.log)-1) < in.GetPrevLogIndex() || s.log[in.GetPrevLogIndex()].Term != in.GetPrevLogTerm()) {
		return &pb.ResponseAppend{Term: s.currentTerm, Success: false}, nil
	}
	s.log = s.log[0 : in.GetPrevLogIndex()+1]
	for i, entry := range in.GetEntries() {
		//TODO append log entry for worker
		log.Printf("Server %v : RequestAppendRPC : Received entry : %v at index %v", serverID, entry, i)
		s.log = append(s.log, entry)
	}
	log.Printf("Server %v : RequestAppendRPC : Received leaderCommit : %v", serverID, in.GetLeaderCommit())
	s.leaderId = in.GetLeaderId()
	if in.GetLeaderCommit() > s.commitIndex {
		s.commitIndex = int64(math.Min(float64(in.GetLeaderCommit()), float64(len(s.log)-1)))
	}
	//TODO commitIndex
	return &pb.ResponseAppend{Term: s.currentTerm, Success: true}, nil
}

func (s *server) AppendRPC(address string, serverID int64) bool {
	// TODO go routine

	leaderId, _ := strconv.Atoi(os.Getenv("CandidateID"))
	leaderID := int64(leaderId)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Server %v : AppendRPC : did not connect: %v", leaderId , err)
		return false
	}
	defer conn.Close()
	c := pb.NewRPCServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	logLength := len(s.log)
	nextLogIndex := s.nextIndex[serverID-1]

	for nextLogIndex >= 0 {
		log.Printf("Server %v : AppendRPC : nextLogIndex : %v", leaderId, nextLogIndex)
		prevLogIndex := nextLogIndex - 1
		log.Printf("Server %v : AppendRPC : prevLogIndex : %v",leaderId, prevLogIndex)
		var prevLogTerm int64
		if prevLogIndex >= 0 {
			prevLogTerm = s.log[prevLogIndex].Term
		}
		response, err := c.RequestAppendRPC(ctx, &pb.RequestAppend{Term: s.currentTerm, LeaderId: leaderID,
			PrevLogIndex: prevLogIndex, PrevLogTerm: prevLogTerm,
			Entries: s.log[nextLogIndex:], LeaderCommit: s.commitIndex})
		if err != nil {
			log.Fatalf("Server %v : AppendRPC : did not connect: %v", leaderId , err)
			return false
		}
		log.Printf("Server %v : AppendRPC : Response Received : %s", serverID, response.String())
		if !response.GetSuccess() {
			if response.GetTerm() > s.currentTerm {
				// WHAT TO DO WHEN FOLLOWER'S TERM IS HIGHER THAN LEADER?
				return false
			} else {
				//try again
				nextLogIndex--
				s.nextIndex[serverID-1] = nextLogIndex
			}
		} else {
			s.nextIndex[serverID-1] = int64(logLength)
			//TODO update match index
			return true
		}
	}
	// TODO update leader data for each worker

	// TODO Update server currentTerm in all responses
	return false
}

func (s* server) HeartBeat() {
    leaderId, _ := strconv.Atoi(os.Getenv("CandidateID"))
    NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	log.Printf("Server %v : HeartBeat : NUMBER OF REPLICAS :%v", leaderId , REPLICAS)
	for {
	    log.Printf("Server %v : ElectionWaitTimer value :%v", leaderId , ElectionWaitTimerReset)
        s.ResetTimer()
        if !ElectionWaitTimerReset {
            for i := 1; i <= REPLICAS; i++ {
                    serverId := strconv.Itoa(i)
			        address := "server" + serverId + ":" + os.Getenv("PORT") + serverId
                    if int64(i) == s.leaderId {
                        continue
                    }
                    log.Printf("Server %v : HeartBeat : Send to Follower : %v",leaderId,i)
                    s.AppendRPC(address,int64(i)) // Need to parallelize this 
            }
        }
        mutex.Lock()
        ElectionWaitTimerReset = false
        mutex.Unlock()
        time.Sleep(time.Duration(ElectionWaitTime) * time.Millisecond)
        mutex.Lock()
        if s.leaderId!=int64(leaderId){
            mutex.Unlock()
            break
            }
        mutex.Unlock()
        }
}
