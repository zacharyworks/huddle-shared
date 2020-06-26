package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// DbCon connection
var DbCon *sql.DB

// ConnectDB Connects to database
func ConnectDB() {
	var err error
	DbCon, err = sql.Open("mysql",
		"monty:python@tcp(127.0.0.1:3306)/huddle")

	if err != nil {
		log.Fatal(err)
	} else {
		println("connected")
	}

}
