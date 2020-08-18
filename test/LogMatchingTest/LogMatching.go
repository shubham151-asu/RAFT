package main

import (
	"context"
	"log"
	"time"
	"strconv"
	pb "../../server/gRPC"
	"google.golang.org/grpc"
)



var portPrefix string = "localhost:5000"

func LogMatching(serverId int)([]*pb.LogsResponseLogEntry){
    serverID := strconv.Itoa(serverId)
    log.Printf("Getting Server %v logs",serverID)
    message := pb.LogsRequest{ReportLog:"Give me Logs"}
    address := portPrefix + serverID

    conn, err := grpc.Dial(address , grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Unable to make connection")
    }
    c := pb.NewRPCServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    response, err := c.LogRequestRPC(ctx,&message)
    if err != nil {
	     log.Printf("Error not nil")
    }
    return response.GetEntries()
}
