package main

import "fmt"

func main() {
    // Create a channel for messages
    messages := make(chan string)

    // Spawn a goroutine to send a message and close the channel
    go func() {
        messages <- "Hello" // Send a value to the channel
        close(messages)     // Close the channel
    }()

    // Loop to receive messages from the channel
    for msg := range messages {
        fmt.Println(msg) // Receive values from the channel
    }
}