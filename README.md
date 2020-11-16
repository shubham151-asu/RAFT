# Summary 
 The project is an implementation of RAFT consensus algorithm based on raft extended paper
 and learnings from MIT Lecture on Distributed Systems. A fault-tolerant distributed Key-Value
 store system has been built and features such as consistency and high availability have been
 achieved as a part of RAFT implementation. 
 
# Distributed Systems 
 Collection of autonomous computers that appear to the user as one integrated system. Overall
 these computers give an abstraction of a purpose that they are built for. E:g A Distributed
 File Systems such as GFS is a file system for a user or application but internally 
 it distributes the file data into chunks to provide high performance for certain R/W tasks
 
 In this project, the distributed system developed provides a fuctionality of a key-Value store
 while abstracts the non-functional feature of consistency, reliablity and high availability. 

# RAFT
 To develop a distributed systems that stores key-value pair service which is high available, 
 we haveused RAFT consensus algorithm. RAFT replicates the state Machine (Key-Value pair system in this case)
 onto multiple machines. If the states are replicated and are consistent, the quorom cluster
 would responds even in the event of a server crash . RAFT also guarantees the writes are consistent 
 through its implementations. 
  
# How does RAFT work
 Time Synchronization :
 For any distributed system, one of the most important challenge is time sync. Different DS
 architecture have different time sync implementation such as two-time scheme using lamport's
 global clock in SPANNER. However, the RAFT uses term for the syncing time. For any request 
 response sent by one raft system to another term is shared to verify if the request is valid 
 and also to update times
  
  
  
 Internal State of RAFT replicated server
 All machines/server in a RAFT maintains some state of their own. 
  
 - Persistent States
     - CurrentTerm : Every server maintains a current state or type of global clock that each server
       updates if its lagging behind
     - Logs : All request received by the leader is persistently stored first by the leader in its database
       and leader sends RPCs in parallel to follower so that they commit the logs in their state. A typical 
       data structure for log is logIndex,LogTerm,LogData(commands,key and values)
     - votedFor : All servers maintain a votedFor information persistently to remember which candidate they
       voted in any given term
 - Volatile States
     - PrevLogIndex : Index of last log request send by leader
     - PrevLogTerm : Term for last term
     - CommitIndex : Index at which server has saved the state into State Machine
     - Server current State : All server can take any given role a Leader,worker, and Follower. Based on the
       state, the server does following tasks like sending RPCs or certain type of RPCs.
         - Leader : To maintain consistency, raft follows master-worker or leader-follower approach and only
           leader decide the course of action for any command. All client request are either received by leader or
           redirected to leader from followers. On reception of request leader appends logs to all followers and
           waits until majority(more than half) of them responds. Once majority response is received, leader responds to the client.
           For any given term, it is guaranteed that only one leader will be elected. In case, leader crashes
           re-election happens and some candidate which was follower is chosen as a leader.
         
         - Follower
           Only receives request from the leader and if the requests are valid would respond to requests of the leader
           Initially all server start with a follower state.
           
         - Candidate :
           On expiration of election time, a follower would change its state to candidate and send RPCs by increasing its currentterm
           by 1. If the candidate gets majority votes, election is won and candidate becomes a leader and handles 
           all client requests
     - NextLogIndex : An array to store indexes for each follower its next logIndex to be send to that 
       follower. In ideal case, all nextLogIndex for all follower remain same but they occasionally defer
       if some follower is lagging behind due to crash or drop in message. (initialized to leader log index + 1)
     - MatchIndex : An array to store indexes for each follower its commitIndex or logIndex which the leader
       knows has been added to the state Machine. (initialized to zero, increases monotonically)
    
 - ElectionTime
     To ensure RAFT system works, there should be a leader for each term to respond to client request
     and ensure replication on followers. In the events of crash or network down or packets drop from/of leader
     the followers cannot receive the RPCs. In that event every server maintains a Election timer that would
     expire if leader does not send any request. Such a follower whose timer gets expired first would change
     its state from follower to candidate, vote itelf, and send voting request by using an RPC and wait for majority
     response. If majority votes are received the candidate becomes a leader and all client request are handled
     by the newly elected leader. If majority votes are not received the candidate changes it's state back to follower
     and some other follower whose timer is expired next will start an election by following same method.
 
 
 - Message passing : Any distributed system need to get request/response for any event that
     needs to informed to peer systems. This allows the systems to proceed in a safe manner, while
     the system move forward by servicing the client request. RAFT uses rpc for message passing
     In this project, google protocol buffer is used which is fast compared to REST. 
     RPCs are mode of message passing from one machine to another using Internet protocols like TCP/IP or other
     RPC enable one system to run a procedure on other system remotely. 
 
     - ClientRequestRPC : Clients use this RPC to use distributed key-value store service. Clients send commands like Get or Put
       with values (for put calls) and expect a return for Get calls (most recent writes).  
     
     - RequestAppendRPC : This RPC is used by leader to get majority response from all followers for any 
       client request. This helps leader ensure to maintain an order of all events that are received from clients
       to maintain consistency and drive the system forward. Leader sends logs "entries", "prevLogIndex","PrevLogTerm", 
       "leaderCommit", and "leaderId" to all its followers in AE RPC. Upon reception of request, follower validates
       each request. For a valid request log entries are committed by the follower and positive response is sent. 
       For an invalid request, a negative response is sent.
     
     - RequestVoteRPC : This RPC is used by a candidate to elect itself as leader upon getting majority vote 
       response. All servers maintain a persistent information about votedFor for any given term and ensure that
       they do not vote to two different candidate in the same term

    For all RPCs send by leader or candidate, has previousLogIndex and previousLogTerm. This ensures that 
    the request is valid if the both of the parameters match. 
 
 - WorkFlow of how raft handles client request
        
      Time&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Leader&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Follower1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Follower2 <br />
      &nbsp;|&nbsp;Client |----CR----->|                                         <br />
      &nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|(LC)|------AE1----->|---------AE2------->|    <br />
      &nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;      |<--------RAE2--------|(LC)                      <br />
      &nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|<---------------------RAE2-------------------|(LC) <br />
      &nbsp;|&nbsp;Client |<---RCR---|                                          <br />
          \\/                                                                 <br />
  - Acronyms :
       - CR : ClientRequestRPC
       - AE : AppendEntryRPC
       - RAE : Response for AppendEntryRPC
       - RCR : Response for Client Request RPC
       - LC : Log commit by that server
  - ClientRequestRPC(CR) send by a certain client, is either received by leader or redirected by a follower to the leader
    Upon receiving request, leader commits logs(LC) into its database and sends AppendEntry(RAE) request in parallel
    to all followers(See data structure in Request RPC section ). Upon receiving the request follower validates
    the request, if the request is valid commit log(s) and responds to the leader (RAE). Upon reception of majority
    of response leader responds to the client that the request is complete.
    <br />Note : Majority is at least 1 more than half, for the timing chart shown above majority is 2. Leader always count
    itself while counting majority

# RAFT features and explanation
- RAFT ensures consistency as leader consults majority of replicas before responding to client, which means
  by the time of responding to the client a majority of the replicas have most recent piece of writes
- In an event of crash upto certain failure, raft system stills works as it is not dependent on responses from
  all replicas but majority. It is this a highly available system
- RAFT implementation ensures that any server lagging due to crash or network issues gets as updated as leader
  over the period of time. Thus, raft based systems automatically recover in the event of crash or lag
 
  
# RUN
- To Run RAFT based Key-Value store system
  - A Docker environment has been developed. If 3 REPLICAS are required in configFile use NUMREPLICAS=3 <br />
   use command :docker-compose up --build 
   
- To test
    go inside test folder(s) and run <br />
    go test -v (assuming go installed in the machine)  
  
 
 
 
 
 Link: https://pdos.csail.mit.edu/6.824/papers/raft-extended.pdf
