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
	password = "" //no my pw is not that lol
	dbname   = "postgres"
)

// queries
const (
	CREATE_USER_ALL_CARDS_TABLE = `CREATE TABLE user_all_cards(
		id serial PRIMARY KEY NOT NULL,
		card_name VARCHAR(128) NOT NULL,
		mana_cost VARCHAR(255) NOT NULL,
		mana_cost_colors CHAR[],
		types VARCHAR[] );`
	DROP_USER_ALL_CARDS_TABLE   = `DROP TABLE IF EXISTS user_all_cards;`
	INSERT_USER_ALL_CARDS_TABLE = `INSERT INTO user_all_cards VALUES 
		('Black Lotus', '5', '{'B', 'W'}', '{})`
)

/* for checking errors with connections, I'll decide later if I want to panic
or set some type of retry loop with retry limit
TODO:
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

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	return db, nil
}

func CreateTablesAllCards(db *sql.DB) error {
	if _, err := db.Exec(CREATE_USER_ALL_CARDS_TABLE); err != nil {
		return err
	}
	return nil
}
