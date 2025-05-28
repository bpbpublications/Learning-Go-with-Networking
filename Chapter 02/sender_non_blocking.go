package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to connect:", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := "Hello, Receiver!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Failed to send message:", err)
		os.Exit(1)
	}

	// Simulate some other tasks while waiting for the acknowledgment
	for i := 0; i < 5; i++ {
		fmt.Println("Performing other tasks...")
		time.Sleep(time.Second)
	}

	// Check for acknowledgment
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to receive acknowledgment:", err)
		os.Exit(1)
	}

	ack := string(buffer)
	fmt.Println("Acknowledgment received:", ack)
}
