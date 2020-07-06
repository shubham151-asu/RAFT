package RPCs

import
(
	"log"
    "os"
    "strconv"
     pb "raftAlgo.com/service/server/gRPC"
	"context"
)


func (s *server) ClientRequestRPC(ctx context.Context, in *pb.ClientRequest) (*pb.ClientResponse, error) {
	log.Printf("Received term : %v", in.GetCommand())
	//TODO if leader else redirect,  append to log entry for leader
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	log.Printf("NUMBER OF REPLICAS :%v", REPLICAS)
	// Here All server object data will be manipulated
	for i := 1 ; i<=REPLICAS ; i++ {
		serverId := strconv.Itoa(i)
		address := "server"+ serverId + ":"+os.Getenv("PORT")+serverId
		log.Printf("Address of the server : %v",address)
		//TODO Check Leader Status and run AppendEntry threads for all clients and count number of successfull APE
		s.AppendRPC(in.GetCommand(),address)
		}	
	return &pb.ClientResponse{Success:true, Result:"a,b added"}, nil // dummy return Ensure code should not come here
}
