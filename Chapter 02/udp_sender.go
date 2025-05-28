package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    // Server address
    addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
    if err != nil {
        fmt.Println("Failed to resolve server address:", err)
        return
    }

    // Create a UDP connection
    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        fmt.Println("Failed to connect to the server:", err)
        return
    }
    defer conn.Close()

    // Send three messages
    sendMessage(conn, "Message 1")
    sendMessageWithDelay(conn, "Message 2", time.Second) // Introduce a delay of 1 second
    sendMessage(conn, "Message 3")

    fmt.Println("All messages sent.")
}

func sendMessage(conn *net.UDPConn, message string) {
    // Convert the message to bytes
    data := []byte(message)

    // Send the message
    _, err := conn.Write(data)
    if err != nil {
        fmt.Println("Failed to send message:", err)
        return
    }

    fmt.Println("Sent:", message)
}

func sendMessageWithDelay(conn *net.UDPConn, message string, delay time.Duration) {
    // Convert the message to bytes
    data := []byte(message)

    // Wait for the specified delay
    time.Sleep(delay)

    // Send the message
    _, err := conn.Write(data)
    if err != nil {
        fmt.Println("Failed to send message:", err)
        return
    }

    fmt.Println("Sent with delay:", message)
}