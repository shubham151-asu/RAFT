package DB

import (
 	"database/sql"
 	 "log"
 	 "os"
 	 "strconv"
     _ "github.com/mattn/go-sqlite3"
     pb "raftAlgo.com/service/server/gRPC"
)

type Conn struct {
   DB *sql.DB
}



func (db *Conn) DBInit(){
    serverId := os.Getenv("CandidateID")
    log.Printf("Server %v : DBInit : Initializing Database ", serverId)
	DB , err := sql.Open("sqlite3", "./logger.db?_journal_mode=WAL&_synchronous=NORMAL") //TODO rename db
 	if err != nil {
		log.Fatal("Server %v : DBInit : unable to Database ", serverId)
 	}
 	db.DB = DB
}

func (db *Conn) CreateDBStructure() (lastLogIndex,lastLogTerm int64){
    serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : CreateDBStructure : Creating table 'logs'", serverId)
	statement, err := db.DB.Prepare("CREATE TABLE IF NOT EXISTS logs (logIndex INTEGER PRIMARY KEY, term INTEGER, command TEXT)")
	if err != nil {
		log.Fatal("err : ", err)
	}

	statement.Exec()
	log.Printf("Server %v : CreateDBStructure : Creating Table 'common'", serverId)
	statement, err = db.DB.Prepare("CREATE TABLE IF NOT EXISTS common (property TEXT PRIMARY KEY, value INTEGER)")
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement.Exec()

	rows, err := db.DB.Query("SELECT property, value FROM common")
	if err != nil {
		log.Fatal("err : ", err)
	}
	if !rows.Next() {
		log.Printf("Server %v : CreateDBStructure : Created Table 'common'", serverId)
		statement, err = db.DB.Prepare("INSERT INTO common (property, value) VALUES (?, ?)")
		if err != nil {
			log.Fatal("err : ", err)
		}
		statement.Exec("currentTerm", 0)
		statement.Exec("votedFor", 0)
	} else {
	    log.Printf("Server %v : CreateDBStructure : Table 'common' exists", serverId)
	}
	rows, err = db.DB.Query("SELECT logIndex,term FROM logs WHERE logIndex = (SELECT MAX(logIndex) FROM logs)")
	if err != nil {
		log.Fatal("err : ", err)
	}
	lastLogIndex = -1
	if rows.Next() {
		log.Printf("Server %v : CreateDBStructure : Table 'logs' exists", serverId)
		var logIndex int
		var term int
		rows.Scan(&logIndex, &term)
		lastLogIndex = int64(logIndex)
		lastLogTerm = int64(term)

	} else {
		log.Printf("Server %v : CreateDBStructure : Created Table 'logs'", serverId)
	}
	log.Printf("Server %v : CreateDBStructure : lastLogIndex : %v && lastLogTerm : %v",serverId, strconv.Itoa(int(lastLogIndex)),strconv.Itoa(int(lastLogTerm)))
	return lastLogIndex, lastLogTerm
}
func (db *Conn) InsertLog(logIndex1 int, term1 int, command1 string) {
    serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : InsertLog : logIndex : %v &&  term : %v && command : %v ", serverId,strconv.Itoa(int(logIndex1)),strconv.Itoa(int(term1)),command1)
	tx, err := db.DB.Begin()
	log.Printf("Server %v : InsertLog : Transaction Begins",serverId)

	if err != nil {
		log.Fatal("err : ", err)
	}
	statement, err := db.DB.Prepare("INSERT INTO logs (logIndex,term, command) VALUES (?, ?,?)")
	//log.Printf("Server %v : InsertLog : Statement Prepared",serverId)

	if err != nil {
		log.Fatal("err : ", err)
	}
	_, err = tx.Stmt(statement).Exec(logIndex1, term1, command1)
	//log.Printf("Server %v : InsertLog : Statement Executed",serverId)

	if err != nil {
		log.Printf("Server %v : InsertLog : Transaction Rollback",serverId)
		tx.Rollback()
	} else {
		tx.Commit()
		log.Printf("Server %v : InsertLog : Transaction Successfully Committed",serverId)
	}

	// statement, err = s.db.Prepare("SELECT term,command FROM logs WHERE logIndex = ?")
	// rows, err := statement.Query(logIndex1)
	// if err != nil {
	// 	log.Fatal("err in getLog : ", err)
	// }
	// var term int
	// var command string
	// if rows.Next() {
	// 	rows.Scan(&term, &command)
	// 	log.Printf("insertLog --> lastLogTerm:" + strconv.Itoa(term) + "  command: " + command)
	// }
	// log.Printf("insertLog --> return")

	// LastInsertId, err := res.LastInsertId()
	// RowsAffected, err := res.RowsAffected()
	// log.Printf("LastInsertId : %v  RowsAffected: %v", LastInsertId, RowsAffected)
}
func (db *Conn) InsertBatchLog(lastLogIndex int, logList []*pb.RequestAppendLogEntry)(finalLogIndex, term int64){
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : InsertBatchLog : Transaction Begins",serverId)
    tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement, err := db.DB.Prepare("INSERT INTO logs (logIndex,term, command) VALUES (?, ?,?)")
	//log.Printf("Server %v : InsertBatchLog : Statement Prepared",serverId)

	if err != nil {
		log.Fatal("err : ", err)
	}
	term = logList[0].Term
	for _, logEntry := range logList {
		lastLogIndex++
		_, err = tx.Stmt(statement).Exec(lastLogIndex, logEntry.Term, logEntry.Command)
		term = logEntry.Term
		log.Printf("Server %v : InsertBatchLog : lastLogIndex : %v && term : %v && Entry : %v",serverId,lastLogIndex,term,logEntry)
	}
	//log.Printf("Server %v : InsertBatchLog : Statement Executed",serverId)

	if err != nil {
		log.Printf("doing rollback")
		tx.Rollback()
	} else {
		tx.Commit()
		//setLastLog(int64(lastLogIndex), term)
		finalLogIndex = int64(lastLogIndex)
		log.Printf("Server %v : InsertBatchLog : Transaction Successfully Committed",serverId)
	}
	log.Printf("Server %v : InsertBatchLog : finalLogIndex : %v && term : %v",serverId,finalLogIndex,term)
	return finalLogIndex , term
}
func (db *Conn) GetLogList(startLogIndex int, endLogIndex int) ( []*pb.RequestAppendLogEntry) { //inclusive
    serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : GetLogList : logIndex : %v &&  term : %v ", serverId, strconv.Itoa(startLogIndex), strconv.Itoa(endLogIndex))
	statement, err := db.DB.Prepare("SELECT logIndex,term,command FROM logs WHERE logs.logIndex >= ? AND logs.logIndex <= ?")
	rows, err := statement.Query(startLogIndex, endLogIndex)
	var response []*pb.RequestAppendLogEntry
	if err != nil {
		log.Printf("err in getLogList : ", err)
		return response
	}
	for rows.Next() {
		var logIndex int
		var term int
		var command string
		rows.Scan(&logIndex, &term, &command)
		log.Printf("Server %v : InsertLog : logIndex : %v &&  term : %v && command : %v ", serverId,strconv.Itoa(logIndex),strconv.Itoa(term),command)
		response = append(response, &pb.RequestAppendLogEntry{Command: command, Term: int64(term)})
	}
	return response
}

func (db *Conn) GetLog(logIndex int) (*pb.RequestAppendLogEntry) {
    serverId := os.Getenv("CandidateID")
    log.Printf("Server %v : GetLog : logIndex : %v",serverId, strconv.Itoa(logIndex))
	statement, err := db.DB.Prepare("SELECT term,command FROM logs WHERE logIndex = ?")
	rows, err := statement.Query(logIndex)
	if err != nil {
		log.Fatal("err in getLog : ", err)
	}
	var term int
	var command string
	if rows.Next() {
		rows.Scan(&term, &command)
		log.Printf("Server %v : GetLog : term : %v && command : %v ", serverId,strconv.Itoa(term),command)
		return &pb.RequestAppendLogEntry{Command: command, Term: int64(term)}
	}
	return nil
}
