package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User

	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate and process user data securely
	if isValidUser(user) {
		processUserData(user)
	} else {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	// Respond to the request
	fmt.Fprintf(w, "Request processed successfully")
}

func isValidUser(user User) bool {
	// Add proper validation logic based on your application's requirements
	// For example, ensure that the username and role meet certain criteria
	// Here, we'll simply check that the role is either "user" or "admin"
	return user.Role == "user" || user.Role == "admin"
}

func processUserData(user User) {
	// Process user data securely
	fmt.Printf("Processing user: %s\n", user.Username)
	if user.Role == "admin" {
		// Log or perform admin actions securely
		fmt.Println("Performing admin actions...")
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
