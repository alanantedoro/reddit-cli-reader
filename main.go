package main

import (
	"log"
	"net/http"

	"redditcli/cmd"

	"redditcli/server"
)

func main() {
	err := server.Auth()
	if err != nil {
		log.Fatal(err)
	}

	cmd.RootCmd.Execute()

	/*
		Ver cuando iniciar el sv, creo que toda esta
		logica deberia estar en el root.go.
	*/

	http.HandleFunc("/getAuth", server.GetAuth)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "server/thanks.html")
	})

	err = server.StartSv()
	if err != nil {
		log.Fatal(err)
	}
}
