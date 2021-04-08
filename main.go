package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
<<<<<<< HEAD
	db, err := sql.Open("mysql", "root:971216@tcp(127.0.0.1:3306)/postanalyzer")
=======
	fmt.Println("Hello World!")

	db, err := sql.Open("mysql", "root:PW@tcp(127.0.0.1:3306/postanalyzer")
>>>>>>> d80c6b24a2fc5c5f9b3545000e8803e32a39dcac
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
