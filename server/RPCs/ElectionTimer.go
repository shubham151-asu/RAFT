package RPCs

import
(
	"log"
	"time"
	"os"
    "math/rand"
    "sync"
    "strconv"
)

var RTT int = 10000
var MTBF int = 50000
var ElectionWaitTime int = 10000

var waitTime int 
var TimerReset bool = false
var ElectionWaitTimerReset bool = false

var mutex sync.Mutex

func (s *server)ResetTimer(){
    candidateId :=  os.Getenv("CandidateID")
    rand.Seed(time.Now().UnixNano())
    waitTime = RTT + rand.Intn(MTBF)
    log.Printf("Server %v : ResetTimer : Timer Reset with WaitTime : %v",candidateId,waitTime)
    mutex.Lock()
    TimerReset = true
    mutex.Unlock()
}

func (s *server)ElectionPreventionTimer(){
    candidateId :=  os.Getenv("CandidateID")
    log.Printf("Server %v : ElectionPreventionTimer : Timer Reset with WaitTime : %v",candidateId,ElectionWaitTime)
    mutex.Lock()
    ElectionWaitTimerReset = true
    mutex.Unlock()
}

func (s *server) WaitForTimeToExpire(){
    done := make(chan bool)
    candidateId :=  os.Getenv("CandidateID")
    for {
        mutex.Lock()
        TimerReset = false
        mutex.Unlock()
        time.Sleep(time.Duration(waitTime) * time.Millisecond)
        if !TimerReset {
	        log.Printf("Server %v : WaitForTimeToExpire : Timer Expired Starting Election",candidateId)
            go s.StartElection(done)
            <-done
        }
    }
}

func (s *server) StartElection(done chan bool) {
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	candidateId :=  os.Getenv("CandidateID")
	CandidateID, _ := strconv.Atoi(candidateId)
	log.Printf("Server %v : StartElection : NUMBER OF REPLICAS :%v", candidateId, REPLICAS)
	count := 1 // Vote self
	finished := 1 // One vote count due to self
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	log.Printf("Server %v : StartElection : Election Started ",candidateId)
	s.currentTerm += 1
	s.votedFor = int64(CandidateID)
	for i := 1 ; i<=REPLICAS ; i++ {
		serverId := strconv.Itoa(i)
		address := "server"+ serverId + ":"+os.Getenv("PORT")+serverId
		if serverId!=candidateId{
		    log.Printf("Server %v : StartElection : address of destination server %s",candidateId, address)
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
    //time.Sleep(time.Duration(waitTime) * time.Millisecond)
	mu.Lock()
	log.Printf("Server %v : StartElection : Votes Received : %v",candidateId,count)
	for count < (REPLICAS/2) && finished != REPLICAS{
		cond.Wait()
	}
	if count >= (REPLICAS/2) && !TimerReset {
	    s.leaderId = int64(CandidateID)
		log.Printf("Server %v : StartElection : Election for term %v won by %v ",candidateId,s.currentTerm,candidateId)
		go s.HeartBeat()
	} else {
		log.Printf("Server %v : StartElection : Election lost ",candidateId)
	}
	mu.Unlock()
    done <- true
    log.Printf("Server %v : StartElection : Returning from Election ",candidateId)

}

func (s * server) ElectionInit() {
    candidateId :=  os.Getenv("CandidateID")
    log.Printf("Server %v : ElectionInit : Election Initialized",candidateId)
    s.ResetTimer()
    s.WaitForTimeToExpire()
}
