package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id   int
	name string
}

func PrintUserName(db *sql.DB, id int) {
	var user User
	rows, err := db.Query("select id, name from users where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.id, &user.name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(user.name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func PrintUserNamePreparedStatement(db *sql.DB, id int) {
	var user User
	stmt, err := db.Prepare("select id, name from users where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	//user QueryRow when you know there is only one record returned
	err = stmt.QueryRow(1).Scan(&user.id, &user.name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user.name)
}

func main() {

	// Creating Database connection manager
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Fetching Values from Database
	PrintUserName(db, 1)
	PrintUserNamePreparedStatement(db, 1)

	var name string
	err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
