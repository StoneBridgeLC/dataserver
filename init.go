package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Initialize database
func initDB() error {
	var dbErr error

	dbUser := os.Getenv("DBUSER")
	dbPw := os.Getenv("DBPW")
	dbName := os.Getenv("DBNAME")
	dbAddr := os.Getenv("DBADDRESS")

	// for local
	dbUser = "root"
	dbPw = "971216"
	dbAddr = "localhost:3306"

	db, dbErr = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4", dbUser, dbPw, dbAddr, dbName))
	if dbErr != nil {
		return dbErr
	}

	return nil
}
