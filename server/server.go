package server

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

var port = "4196"

var clientID = ""
var state = "test"
var redirectURI = "http://localhost:4196/"
var scopes = "identity,edit,mysubreddits,read,report,submit,subscribe,vote"

func StartSv() error {
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":4196", nil); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func OkHandler(w http.ResponseWriter, r *http.Request) {
	/*
		1. Tomar la respuesta, viene luego de la ruta /, necesitamos al menos el token.
		http://localhost:4196/#access_token=eyJhbGciOiJSUzI1NiIsImtpZCI6IlNIQTI1NjpzS3dsMnlsV0VtMjVmcXhwTU40cWY4MXE2OWFFdWFyMnpLMUdhVGxjdWNZIiwidHlwIjoiSldUIn0.eyJzdWIiOiJ1c2VyIiwiZXhwIjoxNzA4MTI2NDcxLjM3MjQ0NywiaWF0IjoxNzA4MDQwMDcxLjM3MjQ0NywianRpIjoiYlZQN1NqUWNJdHlsd0h6ZzlqcF9lUWFicWFxM253IiwiY2lkIjoiX0dxR0t3em1TUDBLWnRfNHRjZmRvUSIsImxpZCI6InQyXzExYWh3NiIsImFpZCI6InQyXzExYWh3NiIsImxjYSI6MTQ3MzQzNDU3MDY0Mywic2NwIjoiZUp3Y3lqRU9nREFNUTlHN2VPNk5FQU1oR1RLVW9zUkY2dTFST2xsUF9nZHlTdDdoWW1nd2RhTGhHeXoxbFZQQ1ZKMkpWbUhmZDlpbGU5NFJaVmQ3NkZ3NF93QUFBUF9fRmxvYWlBIiwiZmxvIjo3fQ.nMEa7DtZgCYRrLeguLkB8ZPnUWXkTsq1q4VNDVrbSuSXLtUFQ192ZnuFF3NV1XnzM-WbSQVW9VZwPDtLMUTcPVxzrFrfNVCdHuI_LbYqjBysp0CD9bNUANOyy5KBqH6QPdahwjXKtQ1ybPkQMUuTVMIw9ZoFdfLCy6-LgUEwyYBV95IaCqZdNbTL_rGPRuNgj3SNnb9OTNq1lVZAJhGadLUeSdZ8MUwEmrke2De2CcUvdOxd7WQJPDOASwgG8vOT4LcQ2XC7_ISN3lmPuSpioYYhIdQusK0AmBzYsVOrIrR-HkT0VhyRb-ijInHlODIzplNW-CZBWBY1h8VMidgHOQ&token_type=bearer&state=test&expires_in=86400&scope=subscribe+edit+vote+mysubreddits+submit+read+report+identity
		2. Investigar el token para ver si necesitamos guardarlo
		3. Devolver thanks.html
	*/

	path := r.URL.Path
	fmt.Println("path :", path)

	// fmt.Println("url", rawUrl)
	accessToken := r.URL.Query().Get("access_token") // no consigue tomar el token
	// Verificar si se proporcionó un token de acceso
	if accessToken == "" {
		http.Error(w, "Token de acceso no proporcionado en la URL", http.StatusBadRequest)
		return
	}

	fmt.Println("Token de acceso:", accessToken)
	http.Redirect(w, r, "/thanks", http.StatusSeeOther)
}

// Es un gran problema abrir el navegador y manejar la respuesta de reddit desde ahi, quizas lo mejor sea hacer un GET directo al link
// y manejar la respuesta yo sin browser.
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
