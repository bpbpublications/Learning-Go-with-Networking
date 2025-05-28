package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    // Listen for incoming connections
    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Failed to start listener:", err)
        os.Exit(1)
    }
    defer listener.Close()

    fmt.Println("Waiting for incoming connection...")

    // Accept incoming connection
    conn, err := listener.Accept()
    if err != nil {
        fmt.Println("Failed to accept connection:", err)
        os.Exit(1)
    }
    defer conn.Close()

    // Read message from the sender
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Failed to read message:", err)
        os.Exit(1)
    }
    message := string(buffer[:n])
    fmt.Println("Received message:", message)

    // Send acknowledgment to the sender
    ack := "Message received!"
    _, err = conn.Write([]byte(ack))
    if err != nil {
        fmt.Println("Failed to send acknowledgment:", err)
        os.Exit(1)
    }
}