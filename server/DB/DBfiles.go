// package DB

// import (
// 	"database/sql"
// 	"log"
// 	"strconv"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func DBInit() *sql.DB {
// 	db, err := sql.Open("sqlite3", "./logger.db") //TODO rename db
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//defer db.Close()
// 	createDBStructure(db)
// 	return db
// }

// func createDBStructure(db *sql.DB) {
// 	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS logs (logIndex INTEGER PRIMARY KEY, term INTEGER, command TEXT)")
// 	if err != nil {
// 		log.Fatal("err : ", err)
// 	}
// 	statement.Exec()
// 	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS common (property TEXT PRIMARY KEY, value INTEGER)")
// 	if err != nil {
// 		log.Fatal("err : ", err)
// 	}
// 	statement.Exec()
// 	statement, err = db.Prepare("INSERT INTO common (property, value) VALUES (?, ?)")
// 	if err != nil {
// 		log.Fatal("err : ", err)
// 	}
// 	statement.Exec("currentTerm", 0)
// 	statement.Exec("votedFor", 0)
// 	rows, err := db.Query("SELECT property, value FROM common")
// 	if err != nil {
// 		log.Fatal("err : ", err)
// 	}
// 	var property string
// 	var value int
// 	for rows.Next() {
// 		rows.Scan(&property, &value)
// 		log.Printf(property + " " + strconv.Itoa(value))
// 	}
// }
