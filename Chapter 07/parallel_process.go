package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Simulating a network task that requires parallel processing
	networkTasks := []string{"Task1", "Task2", "Task3", "Task4", "Task5"}

	// Use a WaitGroup to wait for all tasks to complete
	var wg sync.WaitGroup

	for _, task := range networkTasks {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			// Simulate processing time
			time.Sleep(time.Second)
			fmt.Printf("Completed: %s\n", t)
		}(task)
	}

	// Wait for all tasks to complete
	wg.Wait()
}
