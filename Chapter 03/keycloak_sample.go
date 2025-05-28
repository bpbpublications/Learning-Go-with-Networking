package main

import (
	"context"
	"log"
	"net/http"

	"github.com/keycloak/keycloak-adapter-for-go/pkg/keycloak"
)

func main() {
	cfg := keycloak.Config{
		BaseURL:       "http://localhost:8080/auth",
		Realm:         "your-realm",
		ClientID:      "your-client-id",
		ClientSecret:  "your-client-secret",
		EnableLogging: true,
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := keycloak.GetToken(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Use the token for authorization or retrieve user information
		log.Println("Access Token:", token.AccessToken)
	})

	http.Handle("/", keycloak.Protect(handler, cfg))

	log.Fatal(http.ListenAndServe(":8081", nil))
}

