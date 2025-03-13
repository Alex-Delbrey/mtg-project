package main

import (
	"database/sql"
	"fmt"
	// "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

/* for checking errors with connections, I'll decide later if I want to panic
or set some type of retry loop with retry limit
TODO:
-properly make getConn() return connection
-make function for sending data (for now just the fiew fields in the table)
-make function to get data from table*/

func getConn() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
