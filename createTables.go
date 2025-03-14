package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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

func GetConn() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	return db, nil
}
