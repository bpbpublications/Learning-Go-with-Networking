package main

import (
    "fmt"
    "runtime"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

func printLetters() {
    for i := 'A'; i <= 'E'; i++ {
        time.Sleep(300 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}

func main() {
    runtime.GOMAXPROCS(2) // Set the number of available cores

    fmt.Println("Start")

    startTime := time.Now()

    go printNumbers()
    go printLetters()

    time.Sleep(3 * time.Second)

    endTime := time.Now()
    elapsedTime := endTime.Sub(startTime)

    fmt.Println("\nElapsed Time:", elapsedTime)

    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)
    fmt.Println("Memory Usage:", memStats.Alloc)
}
