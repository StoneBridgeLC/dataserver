package main

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

// Initialize database
func initDB() error {
	var dbErr error

	dbUser := os.Getenv("DBUSER")
	dbPw := os.Getenv("DBPW")
	dbName := os.Getenv("DBNAME")

	db, dbErr = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(database-1.cd3rtwalq81j.ap-northeast-2.rds.amazonaws.com:3306)/%s", dbUser, dbPw, dbName))
	if dbErr != nil {
		return dbErr
	}

	return nil
}
