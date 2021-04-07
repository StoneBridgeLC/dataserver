package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	fmt.Println("Hello World!")

	db, err := sql.Open("mysql", "root:PW@tcp(127.0.0.1:3306/postanalyzer")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("show tabales")
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
