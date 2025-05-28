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

	// Process user data (insecure deserialization)
	processUserData(user)

	// Respond to the request
	fmt.Fprintf(w, "Request processed successfully")
}

func processUserData(user User) {
	// Insecure deserialization: Assuming user data is trusted without proper validation
	if user.Role == "admin" {
		// Malicious activity: Perform actions only allowed for admins
		fmt.Println("Performing admin actions...")
	}
	// Process other user data
	fmt.Printf("Processing user: %s\n", user.Username)
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
