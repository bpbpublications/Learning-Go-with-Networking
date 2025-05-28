
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/keycloak/keycloak-adapter-for-go/pkg/keycloak"
)

func main() {
	// Keycloak configuration
	cfg := keycloak.Config{
		BaseURL:       "http://localhost:8080/auth",
		Realm:         "your-realm",
		ClientID:      "your-client-id",
		ClientSecret:  "your-client-secret",
		EnableLogging: true,
	}

	// Handler for protected resources
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := keycloak.GetToken(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Use the token for authorization or retrieve user information
		log.Println("Access Token:", token.AccessToken)
		// Your custom logic here

		w.Write([]byte("Welcome to the protected resource!"))
	})

	// Protect the handler with Keycloak
	http.Handle("/", keycloak.Protect(handler, cfg))

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

