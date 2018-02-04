package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

func connect() {
	db, err := sql.Open("mysql", "amy:zhu@tcp(127.0.0.1:3306/monster")
	if err != nil {
		fmt.Print(err)
	}

	// need to close

	//stmt, err := db.Prepare()

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// ...
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	var name string
	err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func getIdForValue(db string, value string) {
	stmt, err := db.Prepare("select id from ? where name=") // change db for consistency
}