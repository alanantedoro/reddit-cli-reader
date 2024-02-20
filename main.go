package main

import (
	"log"
	"net/http"

	"redditcli/cmd"

	"redditcli/server"
)

func main() {

	go func() {
		cmd.RootCmd.Execute()

	}()

	go func() {

		http.HandleFunc("/getAuth", server.GetAuth)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "server/thanks.html")
		})

		err := server.Auth()
		if err != nil {
			log.Fatal(err)
		}

		err = server.StartSv()
		if err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}
