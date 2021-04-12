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
	dbAddr := os.Getenv("DBADDRESS")

	// for local
	dbUser = "root"
	dbPw = "971216"
	dbAddr = "127.0.0.1:3306"

	db, dbErr = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?useUnicode=true&useJDBCCompliantTimezoneShift=true&useLegacyDatetimeCode=false&serverTimezone=UTC", dbUser, dbPw, dbName, dbAddr))
	if dbErr != nil {
		return dbErr
	}

	return nil
}
