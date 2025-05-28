// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	DatabaseUsername string `json:"db_username"`
	DatabasePassword string `json:"db_password"`
}

var config Config

func main() {
	// Load configuration from file
	err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	// Set up a simple HTTP server
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Ensure that only localhost can access sensitive information
	if r.RemoteAddr != "127.0.0.1:8080" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Access the sensitive information from the configuration
	username := config.DatabaseUsername
	password := config.DatabasePassword

	// Simulate using the sensitive information (in a real app, this could be a database connection, etc.)
	response := fmt.Sprintf("Database Credentials: %s:%s", username, password)
	w.Write([]byte(response))
}

func loadConfig(filename string) error {
	// Read the configuration file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into the Config struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	return nil
}
