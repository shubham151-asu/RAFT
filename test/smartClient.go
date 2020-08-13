package main

import (
	"context"
	"log"
	"strconv"
	"time"

	pb "../server/gRPC"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
)

var NUM_REPLICAS int = 3
var portPrefix string = "localhost:5000"

type Client struct {
	conn  *grpc.ClientConn
	cache *cache.Cache
}

func (c *Client) InitializeCache() {
	log.Printf("Initializing cache")
	c.cache = cache.New(5*time.Minute, 10*time.Minute)
	log.Printf("Setting default leader as server : 1")
	c.cache.Set("leader", "1", cache.DefaultExpiration)
	for i := 1; i <= NUM_REPLICAS; i++ {
		workerId := strconv.Itoa(i)
		c.cache.Set(workerId, "ON", cache.DefaultExpiration)
	}
}

func (c *Client) StartClientConnection() {
	leader, _ := c.cache.Get("leader")
	address := portPrefix + leader.(string)
	log.Printf("Leader address : %v", address)
	conn, err := grpc.Dial(portPrefix+leader.(string), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to make connection")
	}
	c.conn = conn
}

func (c *Client) StopClientConnection() {
	c.conn.Close()
	c.cache.Flush()
	c.conn = nil
}

func (c *Client) SetNewLeader(leader string) {
	c.cache.Set("leader", leader, cache.DefaultExpiration)
}

func (client *Client) AssignNewLeader() (ret bool) {
	message := pb.ClientRequest{Command: "Alive", Health: "Alive", Key: "", Value: ""}
	ret = false
	count := 0
	for i := 1; i <= NUM_REPLICAS; i++ {
		serverId := strconv.Itoa(i)
		count += 1
		client.SetNewLeader(serverId)
		client.StartClientConnection()
		c := pb.NewRPCServiceClient(client.conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		response, err := c.ClientRequestRPC(ctx, &message)
		if err != nil {
			//log.Printf("Error not nil")
			count += 1
			client.conn = nil
			cancel()
			continue
		}
		if response.GetLeaderId() > 0 {
			log.Printf("New leader Set as %v", response.GetLeaderId())
			ret = true
			client.SetNewLeader(strconv.Itoa(int(response.GetLeaderId())))
			break
		}
	}
	//log.Printf("Count of server down",count)
	if count > NUM_REPLICAS {
		ret = false
	}
	return ret
}

func (client *Client) SendMessage(message *pb.ClientRequest) (ret *pb.ClientResponse) {

	log.Printf("Sending Message Health : %v Command : %v Key : %v Value : %v", message.GetHealth(), message.GetCommand(), message.GetKey(), message.GetValue())
	if client.conn == nil {
		client.StartClientConnection()
	}
	for {
		leader, _ := client.cache.Get("leader")
		c := pb.NewRPCServiceClient(client.conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		response, err := c.ClientRequestRPC(ctx, message)
		if err != nil {
			cancel()
			client.conn = nil
			log.Printf("Unable to get timely response from server :%v: error message %v", leader.(string), err)
			success := client.AssignNewLeader()
			if !success {
				log.Fatalf("RAFT system down")
			}
		}
		ret = response
		break
	}
	return ret
}

func (c *Client) Put(key, value string) (response *pb.ClientResponse) {
	message := pb.ClientRequest{Command: "put", Health: "", Key: key, Value: value}
	response = c.SendMessage(&message)
	return response
}

func (c *Client) Get(key string) (response *pb.ClientResponse) {
	message := pb.ClientRequest{Command: "get", Health: "", Key: key, Value: ""}
	response = c.SendMessage(&message)
	return response
}
