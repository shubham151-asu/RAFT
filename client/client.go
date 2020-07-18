package main

import (
	"context"
	"log"
	"time"

	pb "../server/gRPC"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50003"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRPCServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.ClientRequestRPC(ctx, &pb.ClientRequest{Command: "add a b"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", response.String())
}
