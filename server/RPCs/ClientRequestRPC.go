package RPCs

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "raftAlgo.com/service/server/gRPC"

	"google.golang.org/grpc"
	pb "raftAlgo.com/service/server/gRPC"
)

func (s *server) ClientRequestRPC(ctx context.Context, in *pb.ClientRequest) (*pb.ClientResponse, error) {
	log.Printf("Received term : %v", in.GetCommand())
	//TODO if leader else redirect,  append to log entry for leader
	if s.leaderId != nil {
		return nil, "Something went wrong, please try again"
	} else if s.serverId == s.leaderId {
		NUMREPLICAS := os.Getenv("NUMREPLICAS")
		REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
		log.Printf("NUMBER OF REPLICAS :%v", REPLICAS)
		// Here All server object data will be manipulated
		for i := 1; i <= REPLICAS; i++ {
			serverId := strconv.Itoa(i)
			address := "server" + serverId + ":" + os.Getenv("PORT") + serverId
			log.Printf("Address of the server : %v", address)
			//TODO Check Leader Status and run AppendEntry threads for all clients and count number of successfull APE
			s.AppendRPC(in.GetCommand(), address)
		}
		return &pb.ClientResponse{Success: true, Result: "a,b added"}, nil // dummy return Ensure code should not come here
	} else {
		//redirect to leader
		address := "server" + s.leaderId + ":" + os.Getenv("PORT") + s.leaderId
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
