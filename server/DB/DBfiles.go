package DB

import (
 	"database/sql"
 	 "log"
 	 "strconv"
     _ "github.com/mattn/go-sqlite3"
     pb "raftAlgo.com/service/server/gRPC"
)

type Conn struct {
   DB *sql.DB
}

func (db *Conn) DBInit(){
	DB , err := sql.Open("sqlite3", "./logger.db?_journal_mode=WAL&_synchronous=NORMAL") //TODO rename db
 	if err != nil {
		log.Fatal(err)
 	}
 	db.DB = DB
}

func (db *Conn) CreateDBStructure() (lastLogIndex,lastLogTerm int64){
	log.Printf("createDBStructure")
	statement, err := db.DB.Prepare("CREATE TABLE IF NOT EXISTS logs (logIndex INTEGER PRIMARY KEY, term INTEGER, command TEXT)")
	if err != nil {
		log.Fatal("err : ", err)
	}
	statement.Exec()
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
		log.Printf("No data in common")
		statement, err = db.DB.Prepare("INSERT INTO common (property, value) VALUES (?, ?)")
		if err != nil {
			log.Fatal("err : ", err)
		}
		statement.Exec("currentTerm", 0)
		statement.Exec("votedFor", 0)
	}
	rows, err = db.DB.Query("SELECT logIndex,term FROM logs WHERE logIndex = (SELECT MAX(logIndex) FROM logs)")
	if err != nil {
		log.Fatal("err : ", err)
	}
	if rows.Next() {
		log.Printf("log exist")
		var logIndex int
		var term int
		rows.Scan(&logIndex, &term)
		lastLogIndex = int64(logIndex)
		lastLogTerm = int64(term)
		log.Printf("lastLogIndex : " + strconv.Itoa(logIndex) + "     lastLogTerm:" + strconv.Itoa(term))
	} else {
		log.Printf("No data in logs")
	}
	return
}
func (db *Conn) InsertLog(logIndex1 int, term1 int, command1 string) {
	log.Printf("insertLog --> logIndex:" + strconv.Itoa(int(logIndex1)) + " term:" + strconv.Itoa(int(term1)) + " Command: " + command1)
	tx, err := db.DB.Begin()
	log.Printf("insertLog --> transaction begin")

	if err != nil {
		log.Fatal("err : ", err)
	}
	statement, err := db.DB.Prepare("INSERT INTO logs (logIndex,term, command) VALUES (?, ?,?)")
	log.Printf("insertLog --> statement Prepared")

	if err != nil {
		log.Fatal("err : ", err)
	}
	_, err = tx.Stmt(statement).Exec(logIndex1, term1, command1)
	log.Printf("insertLog --> statement Exec")

	if err != nil {
		log.Printf("doing rollback")
		tx.Rollback()
	} else {
		tx.Commit()
		log.Printf("insertLog --> transaction commited")
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
	//log.Printf("insertLog --> loglist:" + logList)
	tx, err := db.DB.Begin()
	log.Printf("insertLog --> transaction begin")

	if err != nil {
		log.Fatal("err : ", err)
	}
	statement, err := db.DB.Prepare("INSERT INTO logs (logIndex,term, command) VALUES (?, ?,?)")
	log.Printf("insertLog --> statement Prepared")

	if err != nil {
		log.Fatal("err : ", err)
	}
	term = logList[0].Term
	for _, logEntry := range logList {
		lastLogIndex++
		_, err = tx.Stmt(statement).Exec(lastLogIndex, logEntry.Term, logEntry.Command)
		term = logEntry.Term
		log.Printf("nsertBatchLog --> Exec logEntry : %v at index %v", logEntry, lastLogIndex)
	}
	log.Printf("insertLog --> statement Exec")

	if err != nil {
		log.Printf("doing rollback")
		tx.Rollback()
	} else {
		tx.Commit()
		//setLastLog(int64(lastLogIndex), term)
		finalLogIndex = int64(lastLogIndex)
		log.Printf("insertLog --> transaction commited")
	}
	return
}
func (db *Conn) GetLogList(startLogIndex int, endLogIndex int) ( []*pb.RequestAppendLogEntry) { //inclusive
	log.Printf("getLogList --> startLogIndex:" + strconv.Itoa(startLogIndex) + " endLogIndex:" + strconv.Itoa(endLogIndex))
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
		log.Printf("lastLogIndex : " + strconv.Itoa(logIndex) + "     lastLogTerm:" + strconv.Itoa(term) + "  command: " + command)
		response = append(response, &pb.RequestAppendLogEntry{Command: command, Term: int64(term)})
	}
	return response
}

func (db *Conn) GetLog(logIndex int) (*pb.RequestAppendLogEntry) {
	statement, err := db.DB.Prepare("SELECT term,command FROM logs WHERE logIndex = ?")
	rows, err := statement.Query(logIndex)
	if err != nil {
		log.Fatal("err in getLog : ", err)
	}
	var term int
	var command string
	if rows.Next() {
		rows.Scan(&term, &command)
		log.Printf("getLog --> lastLogTerm:" + strconv.Itoa(term) + "  command: " + command)
		return &pb.RequestAppendLogEntry{Command: command, Term: int64(term)}
	}
	return nil
}
