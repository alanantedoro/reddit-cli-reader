package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

var port = "4196"

var clientID = "SRFxgJ8D1Ydqz_YB66dgUA"
var state = "test"
var redirectURI = "http://localhost:4196/"
var scopes = "identity,edit,mysubreddits,read,report,submit,subscribe,vote"
var authToken = ""

func StartSv() error {
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":4196", nil); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func GetAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	fragment := data["fragment"]
	if fragment == "" {
		http.Error(w, "Fragmento no proporcionado en la solicitud", http.StatusBadRequest)
		return
	}

	// fmt.Println("Fragmento recibido:", fragment)

	index := strings.Index(fragment, "access_token=")
	if index == -1 {
		fmt.Println("No se encontró ningún access token")
		return
	}

	fragment = fragment[index+len("access_token="):]

	endIndex := strings.Index(fragment, "&")
	if endIndex == -1 {
		endIndex = len(fragment)
	}

	accessToken := fragment[:endIndex]

	authToken = accessToken
}

func Auth() error {
	authLink := generateAuthLink()

	err := openBrowser(authLink)

	return err
}

func generateAuthLink() string {
	authLink := fmt.Sprintf("https://www.reddit.com/api/v1/authorize?client_id=%s&response_type=token&state=%s&redirect_uri=%s&scope=%s", clientID, state, redirectURI, scopes)
	return authLink
}

func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
