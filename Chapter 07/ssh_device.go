package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

func main() {
	// Configure the SSH client
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", "192.162.1.1:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// Execute a command
	output, err := session.CombinedOutput("show ip interface brief")
	if err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
	fmt.Println(string(output))
}
