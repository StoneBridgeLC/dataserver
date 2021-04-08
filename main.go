package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:971216@tcp(127.0.0.1:3306)/postanalyzer")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("show tables")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var table string

	for rows.Next() {
		err := rows.Scan(&table)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(table)
	}
}
