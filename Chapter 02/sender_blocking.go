package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
    // Connect to the receiver
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Failed to connect:", err)
        os.Exit(1)
    }
    defer conn.Close()

    // Send message to the receiver
    message := "Hello, Receiver!"
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println("Failed to send message:", err)
        os.Exit(1)
    }

    // Block until the receiver acknowledges the message
    buffer := make([]byte, 1024)
    _, err = conn.Read(buffer)
    if err != nil {
        fmt.Println("Failed to receive acknowledgment:", err)
        os.Exit(1)
    }

    fmt.Println("Message sent and acknowledged:", message)
}