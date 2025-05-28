package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/data", handleDataRequest)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleDataRequest(w http.ResponseWriter, r *http.Request) {
	// Process the request and send the response
	data := "Some data"
	_, err := w.Write([]byte(data))
	if err != nil {
		fmt.Println("Failed to send response:", err)
	}
}
