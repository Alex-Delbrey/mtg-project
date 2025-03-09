package main

import (
	"fmt"
	"net/http"
	// "github.com/gin-gonic/gin"
)

const mtgApi = "https://api.magicthegathering.io/v1/cards"

func main() {
	c, err := http.Get(mtgApi)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c)
}
