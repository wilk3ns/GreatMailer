package main

import (
	"GreatMailer/api"
	"fmt"
)

func main() {
	go api.HandleRequests()

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
