package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	// Data Source Name (username:password@tcp(host:port)/dbname)
	dsn := "yourusername:yourpassword@tcp(yourhost:yourport)/yourdbname" // Replace with your MySQL credentials
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("unable to open the sql err : ", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error : ", err)
		return
	}

	// additional query
	var version string
	err = db.QueryRow("select version()").Scan(&version)
	if err != nil {
		log.Fatal("unable to get the version", err)
		return
	}

	fmt.Println("connected to the database version :", version)
}
