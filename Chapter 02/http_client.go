package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Send a GET request to the server
	resp, err := http.Get("http://localhost:8080/data")
	if err != nil {
		fmt.Println("Failed to send request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		os.Exit(1)
	}

	// Print the response
	fmt.Println("Response:", string(body))
}
