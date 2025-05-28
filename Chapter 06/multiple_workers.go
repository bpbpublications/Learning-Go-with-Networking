package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
    for task := range tasks {
        // Simulate some work
        time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
        result := task * 2
        results <- result
        fmt.Printf("Worker %d processed task %d\n", id, task)
    }
}

func main() {
    numTasks := 10
    numWorkers := 3

    tasks := make(chan int, numTasks)
    results := make(chan int, numTasks)

    // Generate tasks
    for i := 1; i <= numTasks; i++ {
        tasks <- i
    }
    close(tasks)

    var wg sync.WaitGroup

    // Start workers
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            worker(workerID, tasks, results)
        }(i)
    }

    // Collect results
    go func() {
        wg.Wait()
        close(results)
    }()

    // Process results
    for result := range results {
        fmt.Println("Received result:", result)
    }

    fmt.Println("Done")
}