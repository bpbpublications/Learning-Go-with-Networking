package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
    // Create a TCP listener
    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Failed to start the server:", err)
        os.Exit(1)
    }
    defer listener.Close()

    fmt.Println("Server started. Listening on port 8080.")

    // Accept incoming connections
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Failed to accept connection:", err)
            continue
        }

        // Handle connection in a separate Goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Print connection details
    fmt.Println("Accepted connection from:", conn.RemoteAddr())

    // Read data from the connection
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Failed to read data:", err)
        return
    }

    // Convert received bytes to string
    message := string(buffer[:n])
    fmt.Println("Received message:", message)

    // Send response back to the client
    response := "Hello, client!"
    _, err = conn.Write([]byte(response))
    if err != nil {
        fmt.Println("Failed to send response:", err)
        return
    }
    fmt.Println("Response sent:", response)
}