package main

import (
	"fmt"
	"sync"
	"time"
)

func asyncTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("Async Task %d completed.\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go asyncTask(i, &wg)
	}

	wg.Wait()
}
