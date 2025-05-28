package main

import (
	"log"
	"net/http"
)

func main() {
	// Set up a simple HTTP server
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request
	log.Printf("Received request from %s for %s", r.RemoteAddr, r.URL.Path)

	// Your application logic goes here

	// Send a response to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
