package RPCs

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"time"
	pb "raftAlgo.com/service/server/gRPC"
	"google.golang.org/grpc"
)

func (s *server) ClientRequestRPC(ctx context.Context, in *pb.ClientRequest) (*pb.ClientResponse, error) {
	log.Printf("Received Command : %v", in.GetCommand())
	log.Printf("s.leaderId :%v   s.serverId :%v", s.leaderId, s.serverId)
	if s.leaderId == 0 {
		return nil, errors.New("Something went wrong, please try again")
	} else if s.serverId == s.leaderId {
		NUMREPLICAS := os.Getenv("NUMREPLICAS")
		REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
		log.Printf("NUMBER OF REPLICAS :%v", REPLICAS)
		// Here All server object data will be manipulated
		s.log = append(s.log, &pb.RequestAppendLogEntry{Command: in.GetCommand(), Term: s.currentTerm})
		successCount := 1
		for i := 1; i <= REPLICAS; i++ {
			serverId := strconv.Itoa(i)
			address := "server" + serverId + ":" + os.Getenv("PORT") + serverId
			log.Printf("Address of the server : %v", address)
			if int64(i) == s.leaderId {
				continue
			}
			log.Printf("lenth of log : %v", len(s.log))
			//TODO Check Leader Status and run AppendEntry threads for all clients and count number of successfull APE
			if int64(len(s.log)) > s.nextIndex[i-1] {
				log.Printf("calling apend")
				if s.AppendRPC(address, int64(i)) {
					successCount++
				}
				// TODO wait for >50% success response
			}
		}
		if successCount > REPLICAS/2 {
			s.commitIndex++
			return &pb.ClientResponse{Success: true, Result: "a,b added"}, nil // dummy return Ensure code should not come here
		} else {
			return &pb.ClientResponse{Success: false, Result: "a,b was not added"}, nil
		}
	} else {
		//redirect to leader
		address := "server" + strconv.FormatInt(s.leaderId, 10) + ":" + os.Getenv("PORT") + strconv.FormatInt(s.leaderId, 10)
		log.Printf("Address of the server : %v", address)
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewRPCServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		response, err := c.ClientRequestRPC(ctx, &pb.ClientRequest{Command: in.GetCommand()})
		if err != nil {
			log.Fatalf("could not redirect: %v", err)
		}
		log.Printf("redirected: %s", response.String())
		return response, err
	}
}
