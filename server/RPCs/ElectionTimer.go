package RPCs

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var RTT int = 10000
var MTBF int = 5000
var ElectionWaitTime int = 5000

var waitTime int
var TimerReset bool = false
var ElectionWaitTimerReset bool = false

var mutex sync.Mutex

func (s *server) ResetTimer() {
	candidateId := os.Getenv("CandidateID")
	rand.Seed(time.Now().UnixNano())
	waitTime = RTT + rand.Intn(MTBF)
	log.Printf("Server %v : ResetTimer : Timer Reset with WaitTime : %v", candidateId, waitTime)
	s.Lock.Lock()
	TimerReset = true
	s.Lock.Unlock()
}

func (s *server) HeartBeatTimer() {
	candidateId := os.Getenv("CandidateID")
	log.Printf("Server %v : HeartBeatTimer : Timer Reset with WaitTime : %v", candidateId, ElectionWaitTime)
	s.Lock.Lock()
	ElectionWaitTimerReset = true
	s.Lock.Unlock()
}

func (s *server) WaitForTimeToExpire() {
	done := make(chan bool)
	candidateId := os.Getenv("CandidateID")
	for {
		s.Lock.Lock()
		TimerReset = false
		s.Lock.Unlock()
		time.Sleep(time.Duration(waitTime) * time.Millisecond)
		log.Printf("Server %v : WaitForTimeToExpire : Current State : %v", candidateId, s.getState())
		if !TimerReset && s.getState() == follower {
			log.Printf("Server %v : WaitForTimeToExpire : Timer Expired ", candidateId)
			if s.initCandidateDS() {
				log.Printf("Server %v : WaitForTimeToExpire : Candidate Status Granted : %v : Starting Election ", candidateId, s.getState())
				go s.StartElection(done)
				<-done
			} else {
				log.Printf("Server %v : WaitForTimeToExpire : Unable to Set Candidate Status ", candidateId)
			}
		}
	}
}

func (s *server) StartElection(done chan bool) {
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	candidateId := os.Getenv("CandidateID")
	CandidateID, _ := strconv.Atoi(candidateId)
	log.Printf("Server %v : StartElection : NUMBER OF REPLICAS :%v", candidateId, REPLICAS)
	count := 1    // Vote self
	finished := 1 // One vote count due to self
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	log.Printf("Server %v : StartElection : Election Started ", candidateId)
	s.setCurrentTerm(s.getCurrentTerm() + 1)
	s.votedFor = int64(CandidateID)
	log.Printf("Server %v : StartElection : Current State : %v", candidateId, s.getState())
	if s.getState() == candidate {
		for i := 1; i <= REPLICAS; i++ {
			serverId := strconv.Itoa(i)
			address := "server" + serverId + ":" + os.Getenv("PORT") + serverId
			if serverId != candidateId {
				log.Printf("Server %v : StartElection : address of destination server %s", candidateId, address)
				go func(address string) {
					vote := s.VoteRPC(address)
					mu.Lock()
					defer mu.Unlock()
					if vote {
						count++
					}
					finished++
					cond.Broadcast()
				}(address)
			}
		}

		mu.Lock()
		for count < ((REPLICAS/2)+1) && finished != REPLICAS {
			cond.Wait()
		}
		log.Printf("Server %v : StartElection : Votes Received : %v", candidateId, count)
		if count >= ((REPLICAS/2)+1) && !TimerReset && s.getState() == candidate {
			log.Printf("Server %v : StartElection : Election for Term : %v won by Server : %v ", candidateId, s.getCurrentTerm(), candidateId)
			if s.initLeaderDS() {
				log.Printf("Server %v : StartElection : Leader Status Granted : %v Sending HeartBeat ", candidateId, s.getState())
				go s.HeartBeat()
			} else {
				log.Printf("Server %v : StartElection : Unable to start Heartbeat", candidateId)
			}
		} else {
			log.Printf("Server %v : StartElection : Election lost ", candidateId)
		}
		mu.Unlock()
	} else {
		log.Printf("Server %v : StartElection : No Longer a Candidate State : Current State : %v ", candidateId, s.getState())
	}
	done <- true
}

func (s *server) ElectionInit() {
	candidateId := os.Getenv("CandidateID")
	log.Printf("Server %v : ElectionInit : Election Initialized", candidateId)
	s.ResetTimer()
	s.WaitForTimeToExpire()
}
