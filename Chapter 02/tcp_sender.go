package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    for {
        sendMessage()
    }
}

func sendMessage() {
    // Connect to the server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Failed to connect to the server:", err)
        os.Exit(1)
    }
    defer conn.Close()

    // Print connection details
    fmt.Println("Connected to server:", conn.RemoteAddr())

    // Read message from user
    fmt.Print("Enter message: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    message := scanner.Text()

    // Send the message to the server
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println("Failed to send message:", err)
        return
    }

    // Receive response from the server
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Failed to receive response:", err)
        return
    }
    response := string(buffer[:n])
    fmt.Println("Received response:", response)
    fmt.Println("---------------------")
}