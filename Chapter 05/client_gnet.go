package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	// Connect to the echo server on port 9000
	addr := "127.0.0.1:9000"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to echo server. Type 'exit' to quit.")

	// Catch "Ctrl+C" signal to handle cleanup
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nExiting client.")
		conn.Close() // Close the client connection before exiting
		os.Exit(0)
	}()

	// Read input from the user and send messages to the server
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			break
		}

		// Send the message to the server
		_, err := conn.Write([]byte(text + "\n"))
		if err != nil {
			fmt.Println("Error writing to server:", err)
			break
		}
	}

	fmt.Println("Exiting client.")
}