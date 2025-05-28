package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to start listener:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Waiting for incoming connection...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read message:", err)
		return
	}

	message := string(buffer[:n])
	fmt.Println("Received message:", message)

	ack := "Message received!"
	_, err = conn.Write([]byte(ack))
	if err != nil {
		fmt.Println("Failed to send acknowledgment:", err)
		return
	}

	// Simulate some processing time
	time.Sleep(time.Second)

	fmt.Println("Acknowledgment sent:", ack)
}

