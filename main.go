package main

import (
	"fmt"
	"net/http"
	// "github.com/gin-gonic/gin"
)

const mtgApi = "https://api.magicthegathering.io/v1/cards?name=%s"

func main() {
	c, err := http.Get(fmt.Sprintf(mtgApi, "Ancestor's Chosen"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Body.Close()
	fmt.Println(c)
}
