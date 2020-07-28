package RPCs

import (
	"context"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
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
	log.Printf("Server %v : RequestAppendRPC : Received GetEntries : %v", serverID, in.GetEntries())
	// TODO reset timer

	term := in.GetTerm()
	if in.GetTerm() < s.getCurrentTerm() {
		return &pb.ResponseAppend{Term: s.getCurrentTerm(), Success: false}, nil
	}
	lastLogIndex, lastLogTerm := s.getLastLog()
	log.Printf("Server %v : RequestAppendRPC : Length of log : %v", serverID, lastLogIndex+1) // Need to add protection here
	if in.GetPrevLogIndex() >= 0 && (lastLogIndex < in.GetPrevLogIndex() || lastLogTerm != in.GetPrevLogTerm()) {
		return &pb.ResponseAppend{Term: s.currentTerm, Success: false}, nil
	}
	s.ResetTimer()       // Once correct has been verified : Reset your Election Timer
	s.initFollowerDS()   // Once correct term has been verified : Go to Follower State no Matter What was previous State was
	s.currentTerm = term // Updating currentTerm to what sent by leader
	//s.log = s.log[0 : in.GetPrevLogIndex()+1] // Need to add protection here
	// for i, entry := range in.GetEntries() {
	// 	//TODO append log entry for worker
	// 	log.Printf("Server %v : RequestAppendRPC : Received entry : %v at index %v", serverID, entry, i)
	// 	//s.log = append(s.log, entry) // Here as well
	// 	log.Printf("term : %v    command : %v", entry.Term, entry.Command)
	// 	lastLogIndex++
	// 	s.insertLog(int(lastLogIndex), int(entry.Term), entry.Command)
	// 	s.setLastLog(lastLogIndex, entry.Term)
	// }
	if len(in.GetEntries()) > 0 {
		lastLogIndex, lastLogTerm = s.db.InsertBatchLog(lastLogIndex, in.GetEntries())
		s.setLastLog(lastLogIndex, lastLogTerm)
	}
	log.Printf("Server %v : RequestAppendRPC : Received leaderCommit : %v", serverID, in.GetLeaderCommit())
	s.leaderId = in.GetLeaderId() // Need to protect this part
	if in.GetLeaderCommit() > s.getCommitIndex() {
		s.setCommitIndex(int64(math.Min(float64(in.GetLeaderCommit()), float64(lastLogIndex)))) // Need to add protection here
	}
	//TODO commitIndex
	return &pb.ResponseAppend{Term: s.currentTerm, Success: true}, nil
}

func (s *server) AppendRPC(address string, serverID int64) bool {
	response := false
	leaderId, _ := strconv.Atoi(os.Getenv("CandidateID"))
	leaderID := int64(leaderId)
	if s.getState() == leader {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Printf("Server %v : AppendRPC : did not connect: %v", leaderId, err)
			return false
		}
		defer conn.Close()
		c := pb.NewRPCServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		tryAgain := false
		//logLength := len(s.log) // Need to add protection here
		nextLogIndex := s.nextIndex[serverID-1]
		lastLogIndex, _ := s.getLastLog()
		entryList := s.db.GetLogList(int(nextLogIndex), int(lastLogIndex))
		for nextLogIndex >= 0 {
			log.Printf("Server %v : AppendRPC : nextLogIndex : %v", leaderId, nextLogIndex)
			prevLogIndex := nextLogIndex - 1
			log.Printf("Server %v : AppendRPC : prevLogIndex : %v", leaderId, prevLogIndex)
			var prevLogTerm int64
			if prevLogIndex >= 0 {
				prevLogTerm = s.db.GetLog(int(prevLogIndex)).Term
			}
			if tryAgain {
				lastLogIndex, _ = s.getLastLog()
				entryList = s.db.GetLogList(int(nextLogIndex), int(lastLogIndex))
			}
			response, err := c.RequestAppendRPC(ctx, &pb.RequestAppend{Term: s.currentTerm, LeaderId: leaderID,
				PrevLogIndex: prevLogIndex, PrevLogTerm: prevLogTerm,
				Entries: entryList, LeaderCommit: s.commitIndex})
			if err != nil {
				log.Printf("Server %v : AppendRPC : did not connect: %v", leaderId, err)
				return false
			}
			log.Printf("Server %v : AppendRPC : Response Received from server : %v : %v", leaderId, serverID, response.String())
			if !response.GetSuccess() {
				log.Printf("Server %v : AppendRPC : Attempt Failed ", leaderId)
				if response.GetTerm() > s.currentTerm {
					// WHAT TO DO WHEN FOLLOWER'S TERM IS HIGHER THAN LEADER?
					return false
				} else {
					log.Printf("Server %v : AppendRPC : Attempting to Retry ", leaderId)
					tryAgain = true
					nextLogIndex--
					s.nextIndex[serverID-1] = nextLogIndex
				}
			} else {
				log.Printf("Server %v : AppendRPC : Attempt Success", leaderId)
				s.nextIndex[serverID-1] = int64(lastLogIndex + 1)
				return true
			}
		}
	} else {
		log.Printf("Server %v : AppendRPC : No Longer a leader : Current State", leaderId, s.getState())
	}
	// TODO update leader data for each worker

	// TODO Update server currentTerm in all responses
	return response
}

func (s *server) HeartBeat() {
	leaderId, _ := strconv.Atoi(os.Getenv("CandidateID"))
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	log.Printf("Server %v : HeartBeat : NUMBER OF REPLICAS :%v", leaderId, REPLICAS)
	for {
		log.Printf("Server %v : ElectionWaitTimer value :%v", leaderId, ElectionWaitTimerReset)
		s.ResetTimer()
		count := 1    // Vote self
		finished := 1 // One vote count due to self
		var mu sync.Mutex
		cond := sync.NewCond(&mu)
		log.Printf("Server %v : HeartBeat : Current State : %v", leaderId, s.getState())
		if !ElectionWaitTimerReset && s.getState() == leader {
			for i := 1; i <= REPLICAS; i++ {
				serverId := strconv.Itoa(i)
				address := "server" + serverId + ":" + os.Getenv("PORT") + serverId
				if int64(i) == s.leaderId {
					continue
				}
				log.Printf("Server %v : HeartBeat : Send to Follower : %v", leaderId, i)
				go func(address string, id int64) {
					success := s.AppendRPC(address, id)
					mu.Lock()
					defer mu.Unlock()
					if success {
						count++
					}
					finished++
					cond.Broadcast()
				}(address, int64(i))
			}
			mu.Lock()
			for count < ((REPLICAS/2)+1) && finished != REPLICAS {
				cond.Wait()
			}
			log.Printf("Server %v : HeartBeat : Success Count : %v", leaderId, count)
			if count >= ((REPLICAS/2)+1) && !ElectionWaitTimerReset {
				log.Printf("Server %v : HeartBeat : Sent Successfully ", leaderId)
			}
			mu.Unlock()
		}
		mutex.Lock()
		ElectionWaitTimerReset = false
		mutex.Unlock()
		time.Sleep(time.Duration(ElectionWaitTime) * time.Millisecond)
		mutex.Lock()
		if s.getState() != leader {
			log.Printf("Server %v : HeartBeat : No longer a leader : Step back from HeartBeat ", leaderId)
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}
}
