package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type Credentials struct {
	DbURL string
	DbPort string
	DbName string
	DbUser string
	DbPass string
}
// DbCon connection
var DbCon *sql.DB

// ConnectDB Connects to database
func ConnectDB(cred Credentials) {
	var db strings.Builder
	db.WriteString(cred.DbUser)
	db.WriteString(":")
	db.WriteString(cred.DbPass)
	db.WriteString("@tcp(")
	db.WriteString(cred.DbURL)
	db.WriteString(":")
	db.WriteString(cred.DbPort)
	db.WriteString(")/")
	db.WriteString(cred.DbName)
	var err error
	DbCon, err = sql.Open("mysql", db.String())

	if err != nil {
		log.Fatal(err)
	} else {
		println("connected")
	}

}
