package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const mtgApi = "https://api.magicthegathering.io/v1/cards?name=%s"

func main() {
	fmt.Println("Enter mtg card name you wish to find: ")
	var userCard string
	fmt.Scanln(&userCard)

	var cr cardResponse
	resp, err := http.Get(fmt.Sprintf(mtgApi, userCard))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &cr)
	if err != nil {
		fmt.Println(err)
		return
	}

	cards := cr.Cards

	for _, card := range cards {
		if len(card.Rulings) != 0 {
			fmt.Println(card.Rulings[0].Date)
		}
	}

	db, err := GetConn()
	defer db.Close()
	if err != nil {
		fmt.Println("Connection to db failed")
		return
	}
	fmt.Printf("Connection Successful to %T\n", db)

	err = CreateTables(db)
	if err != nil {
		fmt.Printf("Failed to create table(s): %s\n", err)
		return
	}
	fmt.Println("Created tables successfully")
}
