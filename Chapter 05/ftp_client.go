package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const serverAddress = "localhost:2121"

func main() {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Error connecting to the server: %v\n", err)
		return
	}
	//defer conn.Close()

	// Read the initial server response
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Printf("Error reading server response: %v\n", err)
		return
	}

	fmt.Println(string(response[:n]))

	// Send user credentials (replace with your actual username and password)
	sendCommand(conn, "USER admin")
	sendCommand(conn, "PASS password")

	// List the current directory
	//sendCommand(conn, "LIST .")

	// Upload a file (replace with your actual file path)
	fmt.Printf("STORE Command\n")
	sendCommand(conn, "STOR test.txt")

	// Download a file
	sendCommand(conn, "RETR test.txt")

	// Quit the session
	sendCommand(conn, "QUIT")
}

func sendCommand(conn net.Conn, command string) {
	_, err := conn.Write([]byte(command + "\r\n"))
	if err != nil {
		fmt.Printf("Error sending command: %v\n", err)
		return
	}

	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Println(string(response[:n]))

	if isTransferCommand(command) {
		// Handle file transfer
		handleFileTransfer(conn)
	}
}

func isTransferCommand(command string) bool {
	return command == "STOR" || command == "RETR"
}

func handleFileTransfer(conn net.Conn) {
	file, err := os.Create("transferred_file.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err == io.EOF {
			fmt.Println("Transfer complete.")
			break
		} else if err != nil {
			fmt.Printf("Error reading data: %v\n", err)
			break
		}

		file.Write(buffer[:n])
	}
}
