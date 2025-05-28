package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Simulating a network request that may fail
	url := "http://example.com/api/data"

	// Implement retry with exponential backoff
	retries := 3
	for i := 0; i < retries; i++ {
		err := performNetworkRequest(url)
		if err == nil {
			fmt.Println("Request successful!")
			break
		}
		fmt.Printf("Retry %d: %v\n", i+1, err)
		time.Sleep(time.Duration(1<<uint(i)) * time.Second) // Exponential backoff
	}
}

func performNetworkRequest(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Process the response
	// ...

	return nil
}
