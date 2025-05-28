package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr := "localhost:6000" // Update with the server's address
	message := "Hello, server!"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	fmt.Println("Data sent to server:", message)

	// Read response (optional)
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response from server:", string(buffer[:n]))
}
