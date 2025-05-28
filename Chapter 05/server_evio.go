package main
import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/tidwall/evio"
)

func main() {
	var portNum int
	var numLoops int
	var useUDP bool
	var printTrace bool
	var reusePort bool
	var useStandardLib bool

	flag.IntVar(&portNum, "port", 6000, "server port")
	flag.BoolVar(&useUDP, "udp", false, "use UDP mode")
	flag.BoolVar(&reusePort, "reuseport", false, "enable reuseport (SO_REUSEPORT)")
	flag.BoolVar(&printTrace, "trace", false, "print packets to console")
	flag.IntVar(&numLoops, "loops", 0, "number of loops")
	flag.BoolVar(&useStandardLib, "stdlib", false, "use standard library")
	flag.Parse()

	var eventHandlers evio.Events
	eventHandlers.NumLoops = numLoops
	eventHandlers.Serving = func(srv evio.Server) evio.Action {
		log.Printf("Server starting on port %d (loops: %d)", portNum, srv.NumLoops)
		if reusePort {
			log.Printf("ReusePort enabled")
		}
		if useStandardLib {
			log.Printf("Using Standard Library")
		}
		return evio.None
	}
	eventHandlers.Opened = func(conn evio.Conn) (out []byte, opts evio.Options, action evio.Action) {
		log.Printf("Connection opened from %s", conn.RemoteAddr())
		return
	}
	eventHandlers.Closed = func(conn evio.Conn, err error) (action evio.Action) {
		log.Printf("Connection from %s closed: %s", conn.RemoteAddr(), err)
		return
	}
	eventHandlers.Data = func(conn evio.Conn, data []byte) (out []byte, action evio.Action) {
		if printTrace {
			log.Printf("Received from %s: %s", conn.RemoteAddr(), strings.TrimSpace(string(data)))
		}
		return data, evio.None
	}

	netScheme := "tcp"
	if useUDP {
		netScheme = "udp"
	}
	if useStandardLib {
		netScheme += "-net"
	}
	err := evio.Serve(eventHandlers, fmt.Sprintf("%s://:%d?reuseport=%t", netScheme, portNum, reusePort))
	if err != nil {
		log.Fatal(err)
	}
}
