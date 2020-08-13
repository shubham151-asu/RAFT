package DB

import (
	"database/sql"
	"log"
	"math"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	pb "raftAlgo.com/service/server/gRPC"
)

type Conn struct {
	DB *sql.DB
}

func (db *Conn) DBInit() {
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : DBInit : Initializing Database ", serverId)
	DB, err := sql.Open("sqlite3", "./logger.db?_journal_mode=WAL&_synchronous=NORMAL") //TODO rename db
	if err != nil {
		log.Fatal("Server %v : DBInit : unable to Database ", serverId)
	}
	db.DB = DB
}

func (db *Conn) CreateDBStructure() (lastLogIndex, lastLogTerm int64) {
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : CreateDBStructure : Creating table 'logs'", serverId)
	statement, err := db.DB.Prepare("CREATE TABLE IF NOT EXISTS logs (logIndex INTEGER PRIMARY KEY, term INTEGER, command TEXT, key TEXT, value TEXT)")
	if err != nil {
		log.Fatal("err : ", err)
	}

	statement.Exec()
	log.Printf("Server %v : CreateDBStructure : Creating Table 'common'", serverId)
	statement, err = db.DB.Prepare("CREATE TABLE IF NOT EXISTS common (term INTEGER, votedFor INTEGER)")
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement.Exec()

	rows, err := db.DB.Query("SELECT term, votedFor FROM common")
	if err != nil {
		log.Fatal("err : ", err)
	}
	if rows.Next() {
		log.Printf("Server %v : CreateDBStructure : Table 'common' exists", serverId)
		var votedFor int
		var term int
		rows.Scan(&term, &votedFor)
	} else {
		log.Printf("Server %v : CreateDBStructure : Created Table 'common'", serverId)
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
	log.Printf("Server %v : CreateDBStructure : lastLogIndex : %v && lastLogTerm : %v", serverId, strconv.Itoa(int(lastLogIndex)), strconv.Itoa(int(lastLogTerm)))
	return lastLogIndex, lastLogTerm
}

func (db *Conn) SetTermAndVotedFor(currentTerm, votedFor int) {
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : SetTermAndVotedFor : term : %v &&  votedFor : %v ", serverId, currentTerm, votedFor)
	tx, err := db.DB.Begin()
	log.Printf("Server %v : SetTermAndVotedFor : Transaction Begins", serverId)
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement, err := db.DB.Prepare("INSERT INTO common (term, votedFor) VALUES (?,?)")
	if err != nil {
		log.Fatal("err : ", err)
	}
	_, err = tx.Stmt(statement).Exec(currentTerm, votedFor)
	//log.Printf("Server %v : InsertLog : Statement Executed",serverId)
	if err != nil {
		log.Printf("Server %v : SetTermAndVotedFor : Transaction Rollback", serverId)
		tx.Rollback()
	} else {
		tx.Commit()
		log.Printf("Server %v : SetTermAndVotedFor : Transaction Successfully Committed", serverId)
	}
}

func (db *Conn) GetVotedFor(currentTerm int64) (votedFor int64) {
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : GetVotedFor : term : %v", serverId, currentTerm)
	statement, err := db.DB.Prepare("SELECT votedFor FROM common WHERE common.term = ?")

	rows, err := statement.Query(int(currentTerm))
	if err != nil {
		log.Fatal("err in getLog : ", err)
	}

	if rows.Next() {
		var votedTo int
		rows.Scan(&votedTo)
		log.Printf("Server %v : GetVotedFor : currentTerm : %v  && votedFor : %v", serverId, currentTerm, votedTo)
		votedFor = int64(votedTo)
	}
	return votedFor
}

func (db *Conn) InsertLog(logIndex int, term int, command string, key string, value string) (success bool) {
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : InsertLog : logIndex : %v &&  term : %v && command : %v && key : %v && value : %v ", serverId, strconv.Itoa(int(logIndex)), strconv.Itoa(int(term)), command, key, value)
	tx, err := db.DB.Begin()
	log.Printf("Server %v : InsertLog : Transaction Begins", serverId)

	if err != nil {
		log.Fatal("err : ", err)
	}
	statement, err := db.DB.Prepare("INSERT INTO logs (logIndex,term, command, key, value) VALUES (?,?,?,?,?)")
	//log.Printf("Server %v : InsertLog : Statement Prepared",serverId)

	if err != nil {
		log.Fatal("err : ", err)
	}
	_, err = tx.Stmt(statement).Exec(logIndex, term, command, key, value)
	//log.Printf("Server %v : InsertLog : Statement Executed",serverId)

	if err != nil {
		log.Printf("Server %v : InsertLog : Transaction Rollback", serverId)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		log.Printf("Server %v : InsertLog : Transaction Successfully Committed", serverId)
		return true
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
func (db *Conn) DeleteLogGreaterThanEqual(logIndex int) {
	serverId := os.Getenv("CandidateID")
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement, err := db.DB.Prepare("DELETE FROM logs WHERE logIndex >= ?")
	//log.Printf("Server %v : InsertBatchLog : Statement Prepared",serverId)

	if err != nil {
		log.Fatal("err : ", err)
	}
	_, err = tx.Stmt(statement).Exec(logIndex)
	if err != nil {
		log.Printf("doing rollback")
		tx.Rollback()
	} else {
		tx.Commit()
		log.Printf("Server %v : deleteLogGreaterThanEqual :  Successfully deleted all logs >= %v", serverId, logIndex)
	}

}
func (db *Conn) DeleteLogByLogIndex(logIndex int) {
	serverId := os.Getenv("CandidateID")
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement, err := db.DB.Prepare("DELETE FROM logs WHERE logIndex = ?")
	if err != nil {
		log.Fatal("err : ", err)
	}
	_, err = tx.Stmt(statement).Exec(logIndex)
	if err != nil {
		log.Printf("doing rollback")
		tx.Rollback()
	} else {
		tx.Commit()
		log.Printf("Server %v : deleteLogGreaterThanEqual :  Successfully deleted all logs >= %v", serverId, logIndex)
	}

}
func (db *Conn) InsertBatchLog(lastLogIndex int64, logList []*pb.RequestAppendLogEntry) (finalLogIndex, term int64, success bool) {
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : InsertBatchLog : Transaction Begins", serverId)
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement, err := db.DB.Prepare("INSERT INTO logs (logIndex,term, command, key, value) VALUES (?, ?,?,?,?)")
	//log.Printf("Server %v : InsertBatchLog : Statement Prepared",serverId)

	if err != nil {
		log.Fatal("err : ", err)
	}
	term = logList[0].Term
	entryList := db.GetLogList(int(lastLogIndex+1), math.MaxInt64)
	var entryMap map[int64]*pb.RequestAppendLogEntry
	for _, entry := range entryList {
		entryMap[entry.LogIndex] = entry
	}
	lastLogIndexBeforeCommit := lastLogIndex
	for _, logEntry := range logList {
		lastLogIndex = logEntry.LogIndex
		if val, exist := entryMap[logEntry.LogIndex]; exist {
			if val.Term == logEntry.Term {
				log.Printf("Server %v : InsertBatchLog : already updated : term : %v && Entry : %v", serverId, lastLogIndex, term, logEntry)
				continue
			} else {
				db.DeleteLogGreaterThanEqual(int(logEntry.LogIndex))
				entryMap = make(map[int64]*pb.RequestAppendLogEntry)
			}
		}
		_, err = tx.Stmt(statement).Exec(logEntry.LogIndex, logEntry.Term, logEntry.Command, logEntry.Key, logEntry.Value)
		term = logEntry.Term
		log.Printf("Server %v : InsertBatchLog : lastLogIndex : %v && term : %v && Entry : %v", serverId, lastLogIndex, term, logEntry)
	}
	//log.Printf("Server %v : InsertBatchLog : Statement Executed",serverId)

	if err != nil {
		log.Printf("Server %v : InsertBatchLog : Transaction Unsuccessful : Doing Rollback", serverId)
		finalLogIndex = lastLogIndexBeforeCommit
		tx.Rollback()
		return finalLogIndex, term, false
	} else {
		tx.Commit()
		//setLastLog(int64(lastLogIndex), term)
		finalLogIndex = lastLogIndex
		log.Printf("Server %v : InsertBatchLog : Transaction Successfully Committed", serverId)
	}
	log.Printf("Server %v : InsertBatchLog : finalLogIndex : %v && term : %v", serverId, finalLogIndex, term)
	return finalLogIndex, term, true
}
func (db *Conn) GetLogList(startLogIndex int, endLogIndex int) []*pb.RequestAppendLogEntry { //inclusive
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : GetLogList : startLogIndex : %v &&  endLogIndex : %v ", serverId, strconv.Itoa(startLogIndex), strconv.Itoa(endLogIndex))
	statement, err := db.DB.Prepare("SELECT logIndex,term,command,key,value FROM logs WHERE logs.logIndex >= ? AND logs.logIndex <= ?")
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
		var key string
		var value string
		rows.Scan(&logIndex, &term, &command, &key, &value)
		log.Printf("Server %v : GetLogList : logIndex : %v &&  term : %v && command : %v && key : %v && value : %v ", serverId, strconv.Itoa(logIndex), strconv.Itoa(term), command, key, value)
		response = append(response, &pb.RequestAppendLogEntry{Command: command, Key: key, Value: value, Term: int64(term), LogIndex: int64(logIndex)})
	}
	return response
}

func (db *Conn) GetLog(logIndex int) *pb.RequestAppendLogEntry {
	serverId := os.Getenv("CandidateID")
	log.Printf("Server %v : GetLog : logIndex : %v", serverId, strconv.Itoa(logIndex))
	statement, err := db.DB.Prepare("SELECT logIndex,term,command,key,value FROM logs WHERE logs.logIndex = ?")
	rows, err := statement.Query(logIndex)
	if err != nil {
		log.Fatal("err in getLog : ", err)
	}
	var term int
	var command string
	var key string
	var value string
	if rows.Next() {
		rows.Scan(&logIndex, &term, &command, &key, &value)
		log.Printf("Server %v : GetLog : logIndex : %v &&  term : %v && command : %v && key : %v && value : %v ", serverId, strconv.Itoa(logIndex), strconv.Itoa(term), command, key, value)
		return &pb.RequestAppendLogEntry{Command: command, Key: key, Value: value, Term: int64(term), LogIndex: int64(logIndex)}
	}
	return nil
}
