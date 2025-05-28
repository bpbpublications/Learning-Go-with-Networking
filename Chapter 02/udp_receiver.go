package main

import (
    "fmt"
    "net"
)

func main() {
    pc, err := net.ListenPacket("udp", ":8080")
    if err != nil {
        fmt.Println("Failed to start the server:", err)
        return
    }
    defer pc.Close()

    fmt.Println("Server started. Listening on port 8080.")

    buffer := make([]byte, 1024)
    for {
        n, addr, err := pc.ReadFrom(buffer)
        if err != nil {
            fmt.Println("Failed to receive message:", err)
            return
        }
        msg := string(buffer[:n])
        fmt.Printf("Received message from %s: %s\n", addr.String(), msg)
    }
}