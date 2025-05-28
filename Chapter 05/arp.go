package main

import (
	"fmt"
	"net"
	"time"

	"github.com/mdlayher/arp"
)

func main() {
	targetIP := net.ParseIP("192.168.1.1")
	iface, err := net.InterfaceByName("eth0")
	if err != nil {
		fmt.Println("Failed to get interface:", err)
		return
	}

	c, err := arp.Dial(iface)
	if err != nil {
		fmt.Println("Failed to create ARP connection:", err)
		return
	}
	defer c.Close()

	req := &arp.Packet{
		Operation: arp.OperationRequest,
		SenderIP:  net.IP{192, 168, 1, 2}, // Your IP
		TargetIP:  targetIP,
	}

	err = c.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		fmt.Println("Failed to set read deadline:", err)
		return
	}

	err = c.WriteTo(req, targetIP)
	if err != nil {
		fmt.Println("Failed to send ARP request:", err)
		return
	}

	buf := make([]byte, 28)
	n, _, err := c.ReadFrom(buf)
	if err != nil {
		fmt.Println("Failed to read ARP response:", err)
		return
	}

	resp, err := arp.Parse(buf[:n])
	if err != nil {
		fmt.Println("Failed to parse ARP response:", err)
		return
	}

	if resp.Operation != arp.OperationReply {
		fmt.Println("Received non-reply ARP packet")
		return
	}

	fmt.Printf("Resolved IP %s to MAC %s\n", resp.SenderIP, resp.SenderHardwareAddr)
}

