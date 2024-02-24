package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"redditcli/server"
	"redditcli/tui"
)

func main() {

	go func() {
		// cmd.RootCmd.Execute()
		t := tui.InitTUI()
		err := t.Start()
		if err != nil {
			fmt.Printf("Failed to start: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(t)

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
