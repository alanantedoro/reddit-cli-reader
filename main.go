package main

import (
	"fmt"
	"log"
	"net/http"

	"redditcli/server"

	"github.com/rivo/tview"
)

func main() {

	go func() {
		// cmd.RootCmd.Execute()
		t := tui.startTUI()
		fmt.Println(t)
		box := tview.NewBox().SetBorder(true).SetTitle("Reddit-CLI")
		if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
			panic(err)
		}
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
