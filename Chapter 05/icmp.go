package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <hostname or IP>")
		return
	}

	destination := os.Args[1]

	// Resolve destination IP address
	ipAddr, err := net.ResolveIPAddr("ip", destination)
	if err != nil {
		fmt.Println("Error resolving IP address:", err)
		return
	}

	// Create ICMP socket
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}
	defer conn.Close()

	// Create ICMP message
	seq := 1
	message := []byte("Ping payload")
	icmpMsg := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: seq,
			Data: message,
		},
	}

	// Serialize ICMP message
	icmpBytes, err := icmpMsg.Marshal(nil)
	if err != nil {
		fmt.Println("Error marshaling ICMP message:", err)
		return
	}

	// Send ICMP packet
	startTime := time.Now()
	_, err = conn.WriteTo(icmpBytes, ipAddr)
	if err != nil {
		fmt.Println("Error sending ICMP packet:", err)
		return
	} else {
		fmt.Println("Sending ICMP Packet")
	}

	// Receive ICMP response
	reply := make([]byte, 1500)
	_, _, err = conn.ReadFrom(reply)
	if err != nil {
		fmt.Println("Error receiving ICMP response:", err)
		return
	}
	elapsed := time.Since(startTime)

	// Parse ICMP response
	icmpReply, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), reply)
	if err != nil {
		fmt.Println("Error parsing ICMP response:", err)
		return
	}
	
	fmt.Printf("Received ICMP reply from %s: seq=%d time=%.3f ms\n",
		ipAddr.String(), seq, float64(elapsed.Microseconds())/1000)
}
