package main

import (
	"log"

	"main.go/server"
)

func main() {
	err := server.Auth()
	if err != nil {
		log.Fatal(err)
	}

	err = server.StartSv()
	if err != nil {
		log.Fatal(err)
	}
}
