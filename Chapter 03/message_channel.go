package main

import (
	"fmt"
	"time"
)

func main() {
	messageChannel := make(chan string)

	go func() {
		time.Sleep(time.Second)
		messageChannel <- "Hello from the goroutine!"
	}()

	message := <-messageChannel
	fmt.Println(message)
}
