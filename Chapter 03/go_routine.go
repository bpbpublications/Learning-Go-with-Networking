package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("This is running concurrently.")
	}()
	
	// Allow time for the goroutine to execute
	time.Sleep(time.Second)
}
