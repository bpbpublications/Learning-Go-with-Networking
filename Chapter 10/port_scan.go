package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ScanResult represents the result of a port scan.
type ScanResult struct {
	Port  int
	State string // "open" or "closed"
}

// Worker represents a worker that performs port scanning.
func Worker(target string, ports <-chan int, results chan<- ScanResult, wg *sync.WaitGroup, timeout time.Duration) {
	defer wg.Done()

	for port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			results <- ScanResult{Port: port, State: "closed"}
			continue
		}
		conn.Close()
		results <- ScanResult{Port: port, State: "open"}
	}
}

// PortScanner scans the specified target for open ports in the given range.
func PortScanner(target string, startPort, endPort, numWorkers int, timeout time.Duration) []ScanResult {
	var wg sync.WaitGroup
	ports := make(chan int, 100)
	results := make(chan ScanResult, 100)

	// Create worker pool
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(target, ports, results, &wg, timeout)
	}

	// Enqueue ports to be scanned
	go func() {
		for port := startPort; port <= endPort; port++ {
			ports <- port
		}
		close(ports)
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect scan results
	var scanResults []ScanResult
	for result := range results {
		scanResults = append(scanResults, result)
	}

	return scanResults
}

func main() {
	target := "localhost"
	startPort := 1
	endPort := 1024
	numWorkers := 50
	timeout := 1 * time.Second

	scanResults := PortScanner(target, startPort, endPort, numWorkers, timeout)

	// Display scan results
	for _, result := range scanResults {
		fmt.Printf("Port %d is %s\n", result.Port, result.State)
	}
}
