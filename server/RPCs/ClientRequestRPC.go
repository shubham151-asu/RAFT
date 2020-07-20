package RPCs

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
	"google.golang.org/grpc"
	pb "raftAlgo.com/service/server/gRPC"
)

func (s *server) ClientRequestRPC(ctx context.Context, in *pb.ClientRequest) (*pb.ClientResponse, error) {
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v :  ClientRequestRPC : Received Command : %v", serverId, in.GetCommand())
	log.Printf("Server %v :  ClientRequestRPC : State : %v", serverId, s.getState())
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	if s.leaderId == 0 {
		log.Printf("Server %v :  ClientRequestRPC : No Leader Elected : Come back later", serverId)
		return nil, errors.New("Something went wrong, please try again")
	} else if s.getState() == leader {
		s.HeartBeatTimer()
		log.Printf("Server %v :  ClientRequestRPC : Appending ClientRequest to Logs : Term : %v : Command : %v", serverId, s.getCurrentTerm(), in.GetCommand())
		// Here All server object data will be manipulated
		//s.log = append(s.log, &pb.RequestAppendLogEntry{Command: in.GetCommand(), Term: s.currentTerm}) // Safety
		lastLogIndex, _ := s.getLastLog()
		lastLogIndex++
		//log.Printf("Server %v :  ClientRequestRPC : Incremented lastLogIndex : %v",serverId,lastLogIndex)
		s.db.InsertLog(int(lastLogIndex), int(s.getCurrentTerm()), in.GetCommand())
		s.setLastLog(lastLogIndex, s.getCurrentTerm())
		log.Printf("Server %v :  ClientRequestRPC : length of logs : %v", serverId, lastLogIndex+1)
		count := 1    // Vote self
		finished := 1 // One vote count due to self
		var mu sync.Mutex
		cond := sync.NewCond(&mu)
		log.Printf("Server %v :  ClientRequestRPC : NUMBER OF REPLICAS :%v", serverId, REPLICAS)
		for i := 1; i <= REPLICAS; i++ {
			serverId := strconv.Itoa(i)
			address := "server" + serverId + ":" + os.Getenv("PORT") + serverId
			if int64(i) == s.leaderId {
				continue
			}
			log.Printf("Server %v :  ClientRequestRPC : Address of the server:%v", serverId, address)
			if int64(lastLogIndex) >= s.nextIndex[i-1] && s.getState() == leader {
				log.Printf("Server %v :  ClientRequestRPC : Calling AppendEntry", serverId)
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
		}
		mu.Lock()
		for count < ((REPLICAS/2)+1) && finished != REPLICAS {
			cond.Wait()
		}
		log.Printf("Server %v : ClientRequestRPC : Success Count : %v", serverId, count)
		if count >= ((REPLICAS/2)+1) && s.getState() == leader {
			log.Printf("Server %v : ClientRequestRPC :  Majority Response received : Committing Entry ", serverId)
			s.IncrementCommitIndex(1)                                          // Verify this Increment : Whether one or more than 1
			return &pb.ClientResponse{Success: true, Result: "a,b added"}, nil //
		} else {
			return &pb.ClientResponse{Success: false, Result: "a,b was not added"}, nil
		}
		mu.Unlock()
	} else {
		//redirect to leader
		address := "server" + strconv.FormatInt(s.leaderId, 10) + ":" + os.Getenv("PORT") + strconv.FormatInt(s.leaderId, 10)
		log.Printf("Server %v :  ClientRequestRPC : Redirecting to leader with address :%v", serverId, address)
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		// Need to retry if the servers are busy
		if err != nil {
			log.Printf("Server %v :  ClientRequestRPC : did not connect: %v", serverId, err)
		}
		defer conn.Close()
		c := pb.NewRPCServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		response, err := c.ClientRequestRPC(ctx, &pb.ClientRequest{Command: in.GetCommand()})
		if err != nil {
			log.Printf("Server %v :  ClientRequestRPC : could not redirect: %v", serverId, err)
		}
		log.Printf("Server %v :  ClientRequestRPC : redirected Response %s", serverId, response.String())
		return response, err
	}
	return &pb.ClientResponse{Success: false, Result: "a,b was not added"}, nil // Dummy response
}
