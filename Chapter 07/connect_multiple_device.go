package main

import (
	"fmt"
	"sync"
)

// Device represents a network device
type Device struct {
	Hostname string
	IP       string
}

func gatherData(device Device, wg *sync.WaitGroup, dataChan chan<- string) {
	defer wg.Done()
	// Simulate gathering data from the device
	data := fmt.Sprintf("Data from %s [%s]", device.Hostname, device.IP)
	// Send the data to the main goroutine via the channel
	dataChan <- data
}

func main() {
	devices := []Device{
		{"Device1", "192.168.1.1"},
		{"Device2", "192.168.1.2"},
		// Add more devices as needed
	}

	var wg sync.WaitGroup
	dataChan := make(chan string, len(devices))

	// Start a goroutine for each device
	for _, device := range devices {
		wg.Add(1)
		go gatherData(device, &wg, dataChan)
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(dataChan)
	}()

	// Read data from the channel as it arrives
	for data := range dataChan {
		fmt.Println(data)
	}
}
