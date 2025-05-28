package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
)

// DNS query types
const (
	TYPE_A   = 1
	TYPE_NS  = 2
	TYPE_MD  = 3
	TYPE_MF  = 4
	TYPE_CNAME = 5
	TYPE_SOA = 6
	TYPE_MB  = 7
	TYPE_MG  = 8
	TYPE_MR  = 9
	TYPE_NULL = 10
	TYPE_WKS  = 11
	TYPE_PTR  = 12
	TYPE_HINFO = 13
	TYPE_MINFO = 14
	TYPE_MX   = 15
	TYPE_TXT  = 16
	TYPE_AAAA = 0x1c
)

// DNS header structure
type dnshdr struct {
	ID      uint16
	RD      uint8
	TC      uint8
	AA      uint8
	Opcode  uint8
	QR      uint8
	Rcode   uint8
	Unused  uint8
	Pr      uint8
	Ra      uint8
	QueNum  uint16
	RepNum  uint16
	NumRR   uint16
	NumRRSup uint16
}

// DNS query type names
type typeName struct {
	Type     uint16
	TypeName string
}

var dnsTypeNames = []typeName{
	{TYPE_A, "A"},
	{TYPE_NS, "NS"},
	{TYPE_MD, "MD"},
	{TYPE_MF, "MF"},
	{TYPE_CNAME, "CNAME"},
	{TYPE_SOA, "SOA"},
	{TYPE_MB, "MB"},
	{TYPE_MG, "MG"},
	{TYPE_MR, "MR"},
	{TYPE_NULL, "NULL"},
	{TYPE_WKS, "WKS"},
	{TYPE_PTR, "PTR"},
	{TYPE_HINFO, "HINFO"},
	{TYPE_MINFO, "MINFO"},
	{TYPE_MX, "MX"},
	{TYPE_TXT, "TXT"},
	{TYPE_AAAA, "AAAA"},
}

// Function to get DNS query type from type name
func getType(typeStr string) uint16 {
	for _, t := range dnsTypeNames {
		if dnsTypeEqualFold(typeStr, t.TypeName) {
			return t.Type
		}
	}
	return 0
}

// Case-insensitive string comparison
func dnsTypeEqualFold(a, b string) bool {
	return len(a) == len(b) && strings.EqualFold(a, b)
}

// Calculate Internet checksum for packet
func in_cksum(packet []byte, len int) uint16 {
	var nleft = len
	var sum int
	var w []uint16 = *(*[]uint16)(unsafe.Pointer(&packet))

	for nleft > 1 {
		sum += int(*w)
		w = w[1:]
		nleft -= 2
	}

	if nleft == 1 {
		sum += int(*(*uint8)(unsafe.Pointer(w)))
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum)
}

// Display usage information
func usage(progname string) {
	fmt.Printf("Usage: %s <query_name> <destination_ip> [options]\n"+
		"\tOptions:\n"+
		"\t-t, --type\t\tquery type\n"+
		"\t-s, --source-ip\t\tsource ip\n"+
		"\t-p, --dest-port\t\tdestination port\n"+
		"\t-P, --src-port\t\tsource port\n"+
		"\t-i, --interval\t\tinterval (in millisecond) between two packets\n"+
		"\t-n, --number\t\tnumber of DNS requests to send\n"+
		"\t-r, --random\t\tfake random source IP\n"+
		"\t-D, --daemon\t\trun as daemon\n"+
		"\t-h, --help\n"+
		"\n",
		progname)
}

// Format DNS query name
func nameformat(name, QS string) {
	var bungle, x string
	var elem [128]byte

	QS = ""
	bungle = name
	x = strtok(bungle, ".")
	for x != "" {
		if n := snprintf(elem[:], 128, "%c%s", len(x), x); n == 128 {
			fmt.Println("String overflow.")
			os.Exit(1)
		}
		QS += elem[:]
		x = strtok(nil, ".")
	}
}

// Format reverse DNS query for IP address
func nameformatIP(ip, resu string) {
	var reverse, temp, x string
	var comps [10]string
	var px int

	temp = ip
	reverse = temp
	x = strtok(reverse, ".")
	for x != "" {
		if px >= 10 {
			fmt.Println("Force DUMP:: dumbass, wtf you think this is, IPV6?")
			os.Exit(1)
		}
		comps[px] = x
		px++
		x = strtok(nil, ".")
	}
	for px--; px >= 0; px-- {
		reverse += comps[px] + "."
	}
	reverse += "in-addr.arpa"
	nameformat(reverse, resu)
}

// Create DNS question packet
func makeQuestionPacket(data []byte, name string, qtype uint16) int {
	if qtype == TYPE_A {
		nameformat(name, data)
		binary.BigEndian.PutUint16(data[len(data)+1:], TYPE_A)
	}
	/* for other type query
	if qtype == TYPE_PTR {
		nameformatIP(name, data)
		binary.BigEndian.PutUint16(data[len(data)+1:], TYPE_PTR)
	}
	*/

	binary.BigEndian.PutUint16(data[len(data)+3:], CLASS_INET)

	return len(data) + 5
}

// Read IP addresses from a file (TODO: Implement)
func readIPFromFile(filename string) int {
	// TODO: Implement this function
	return 0
}

func main() {
	var qname string = ""                   
	var qtype uint16 = TYPE_A                 
	var srcIP net.IP = nil                  
	var sinDst net.UDPAddr = net.UDPAddr{}   
	var srcPort uint16 = 0                  
	var dstPort uint16 = 53                  
	var sock int = 0                         
	var number int = 0                       
	var count int = 0                        
	var sleepInterval int = 0                
	var randomIP bool = false                
	var staticIP bool = false                

	const on int = 1

	var packet [2048]byte
	var iphdr *ipv4.Header = (*ipv4.Header)(unsafe.Pointer(&packet))
	var udp *layers.UDP = (*layers.UDP)(unsafe.Pointer(&packet[sizeof(ipv4.Header)]))
	var dnsHeader *dnshdr = (*dnshdr)(unsafe.Pointer(&packet[sizeof(ipv4.Header)+sizeof(layers.UDP)]))
	var dnsData []byte = packet[sizeof(ipv4.Header)+sizeof(layers.UDP)+sizeof(dnshdr):]

	for i := range dnsData {
		dnsData[i] = 0
	}

	for i := range dnsTypeNames {
		dnsTypeNames[i].Type = uint16(i) + 1
	}

	for {
		var dnsDataLen int
		var udpDataLen int
		var ipDataLen int

		if randomIP {
			srcIP = net.IPv4(byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)))
		}

		dnsHeader.ID = uint16(rand.Int())
		dnsDataLen = makeQuestionPacket(dnsData, qname, TYPE_A)

		udpDataLen = sizeof(dnshdr) + dnsDataLen
		ipDataLen = sizeof(layers.UDP) + udpDataLen

		/* update UDP header*/
		if srcPort == 0 {
			udp.SrcPort = layers.UDPPort(rand.Intn(65535))
		}
		udp.DstPort = layers.UDPPort(dstPort)
		udp.Length = uint16(sizeof(layers.UDP) + udpDataLen)
		udp.Checksum = 0

		/* update IP header */
		iphdr.SrcIP = srcIP
		iphdr.DstIP = sinDst.IP
		iphdr.Version = ipv4.Version
		iphdr.IHL = sizeof(ipv4.Header) >> 2
		iphdr.TTL = 245
		iphdr.Protocol = layers.IPProtocolUDP
		iphdr.Length = uint16(sizeof(ipv4.Header) + ipDataLen)
		iphdr.Checksum = 0
		// iphdr.Checksum = in_cksum((char *)iphdr, sizeof(struct ip));

		ret, err := syscall.Write(sock, packet)
		if err != nil {
			// Handle error
			fmt.Printf("sendto error: %v\n", err)
		}

		count++

		if number > 0 && count >= number {
			// done
			break
		}

		if sleepInterval > 0 {
			time.Sleep(time.Duration(sleepInterval) * time.Millisecond)
		}
	}

	fmt.Printf("sent %d DNS requests.\n", count)
}
