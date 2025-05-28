package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- string) {
	time.Sleep(time.Second) // Simulate some work
	ch <- "Hello, Channel!" // Send a value through the channel
}

func receiver(ch <-chan string) {
	msg := <-ch // Receive a value from the channel
	fmt.Println(msg)
}

func main() {
	ch := make(chan string) // Create a channel

	go sender(ch)   // Start the sender goroutine
	go receiver(ch) // Start the receiver goroutine

	time.Sleep(time.Second * 2) // Sleep for a while to allow goroutines to execute
	fmt.Println("Done")
}
