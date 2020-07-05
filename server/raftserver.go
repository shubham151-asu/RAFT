package main

import
(
     "raftAlgo.com/service/server/RPCs"
)


func main() {
	RPCs.RPCInit() // Example to Start RPC port Running : In Future We might have to start this in goroutine
	// TODO Add election timer
	// TODO Add Init for stateMachine
	// TODO Add Init for DB to add LastApplied

}

