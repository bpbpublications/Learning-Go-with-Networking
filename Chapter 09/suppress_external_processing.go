package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Username string `xml:"username"`
	Password string `xml:"password"`
}

func parseXML(r *http.Request) (*User, error) {
	decoder := xml.NewDecoder(r.Body)

	var user User
	err := decoder.Decode(&user)
	if err != nil {
		if err == io.EOF {
			return nil, nil // Ignore EOF errors
		}
		return nil, err
	}
	return &user, nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	user, err := parseXML(r)
	if err != nil {
		http.Error(w, "Invalid XML", http.StatusBadRequest)
		return
	}

	// Process the user object securely
	fmt.Fprintf(w, "Received user: %v", user)
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
