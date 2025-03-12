package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const mtgApi = "https://api.magicthegathering.io/v1/cards?name=%s"

func main() {
	var cr cardResponse
	resp, err := http.Get(fmt.Sprintf(mtgApi, "Lotus"))
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

}
