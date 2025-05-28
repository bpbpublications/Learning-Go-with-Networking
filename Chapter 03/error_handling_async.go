package main

import (
	"errors"
	"fmt"
)

func asyncTaskWithError(errChan chan error) {
	// Simulating an error
	err := errors.New("An error occurred")
	errChan <- err
}

func main() {
	errChan := make(chan error)

	go asyncTaskWithError(errChan)

	select {
	case err := <-errChan:
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task completed successfully.")
		}
	}
}
