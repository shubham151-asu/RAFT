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

var RTT int = 1000
var MTBF int = 5000


var waitTime int 
var TimerReset bool = false

var mutex sync.Mutex


func (s *server)ResetTimer(){
    candidateId :=  os.Getenv("CandidateID")
    waitTime = RTT + rand.Intn(MTBF)
    log.Printf("Server %v : Timer Reset with WaitTime : %v",candidateId,waitTime)
    mutex.Lock()
    TimerReset = true
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
	    log.Printf("Server %v : Timer Expired Starting Election",candidateId)
            s.StartElection(done)
            <-done
        }
    }
}

func (s *server) StartElection(done chan bool) {
	NUMREPLICAS := os.Getenv("NUMREPLICAS")
	REPLICAS, _ := strconv.Atoi(NUMREPLICAS)
	candidateId :=  os.Getenv("CandidateID")
	log.Printf("NUMBER OF REPLICAS :%v", REPLICAS)
	count := 1 // Vote self
	finished := 1 // One vote count due to self
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	log.Printf("Server %v : Election Started ",candidateId)
	for i := 1 ; i<=REPLICAS ; i++ {
		serverId := strconv.Itoa(i)
		address := "server"+ serverId + ":"+os.Getenv("PORT")+serverId
		if serverId!=candidateId{
		    log.Printf("Server %v : address of destination server %s",candidateId, address)
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
        time.Sleep(time.Duration(waitTime) * time.Millisecond)
	mu.Lock()
	log.Printf("Server %v : Votes Received : %v",candidateId,count)
	for count < (REPLICAS/2 + 1) && finished != REPLICAS{
		cond.Wait()
	}
	if count >= (REPLICAS/2 + 1) && !TimerReset {
		log.Printf("Server %v : Election won : New Leader election",candidateId)
	} else {
		log.Printf("Server %v : Election lost ",candidateId)
	}
	mu.Unlock()
    done<-true
    log.Printf("Server %v : Returing from Election ",candidateId)

}

func (s * server) ElectionInit() {
    candidateId :=  os.Getenv("CandidateID")
    log.Printf("Server %v : Election Initialized",candidateId)
    s.ResetTimer()
    s.WaitForTimeToExpire()
}
