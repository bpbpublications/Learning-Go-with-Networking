package main

import (
	"fmt"
	"runtime"
)

const (
	MaxItems = 1000000
)

type Data struct {
	ID   int
	Name string
}

func main() {
	// Set GOMAXPROCS to utilize all available CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create a slice with pre-allocated capacity
	dataSlice := make([]Data, 0, MaxItems)

	// Preallocate memory for all items
	for i := 0; i < MaxItems; i++ {
		data := Data{
			ID:   i,
			Name: fmt.Sprintf("Item %d", i),
		}
		dataSlice = append(dataSlice, data)
	}

	// Use the preallocated data
	for _, data := range dataSlice {
		// Process each item
		processData(data)
	}

	// Release the memory
	dataSlice = nil

	// Run garbage collection explicitly
	runtime.GC()

	// Check memory usage after releasing the memory
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Allocated Memory: %.2f MB\n", float64(memStats.Alloc)/1024/1024)
}

func processData(data Data) {
	// process data
}