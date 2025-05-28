package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"read", "write"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://provider.com/oauth/authorize",
			TokenURL: "https://provider.com/oauth/token",
		},
	}
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != "state" {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := oauthConf.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Access Token: %s", token.AccessToken)
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/callback", handleCallback)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
