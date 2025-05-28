package main

import (
    "fmt"
    "os"
    "os/signal"
    "runtime"
    "runtime/pprof"
    "sync"
    "syscall"
    "time"
)

func printNumbers(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 5; i++ {
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

func printLetters(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 'A'; i <= 'E'; i++ {
        time.Sleep(300 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}

func main() {
    // Create a CPU profiling file
    f, err := os.Create("cpu_profile.prof")
    if err != nil {
        fmt.Println("Failed to create CPU profile:", err)
        return
    }
    defer f.Close()

    // Start CPU profiling
    err = pprof.StartCPUProfile(f)
    if err != nil {
        fmt.Println("Failed to start CPU profile:", err)
        return
    }
    defer pprof.StopCPUProfile()

    fmt.Println("Start")

    var wg sync.WaitGroup
    wg.Add(2)

    go printNumbers(&wg)
    go printLetters(&wg)

    wg.Wait()

    fmt.Println("\nDone")

    // Get runtime information
    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)

    fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())
    fmt.Printf("Allocated Memory: %.2f MB\n", float64(memStats.Alloc)/1024/1024)
    fmt.Printf("Total Memory Allocations: %d\n", memStats.TotalAlloc)
    fmt.Printf("Heap Memory Allocated: %.2f MB\n", float64(memStats.HeapAlloc)/1024/1024)
    fmt.Printf("Heap Memory System: %.2f MB\n", float64(memStats.HeapSys)/1024/1024)

    // Print runtime information
    fmt.Println("Runtime Information:")
    fmt.Println("  Goroutines:", runtime.NumGoroutine())
    fmt.Println("  CPU Usage:", getCpuUsage())
    fmt.Println("  Memory Usage:", getMemoryUsage())
    fmt.Println("  Elapsed Time:", time.Since(startTime))
}

func getCpuUsage() string {
    // Retrieve CPU usage
    elapsed := time.Since(startTime)
    goroutineCount := pprof.Lookup("goroutine").Count()
    cpuUsage := float64(goroutineCount) / elapsed.Seconds()

    return fmt.Sprintf("%.2f%%", cpuUsage*100)
}

func getMemoryUsage() string {
    // Retrieve memory usage
    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)

    return fmt.Sprintf("%.2f MB", float64(memStats.Alloc)/1024/1024)
}

var startTime time.Time

func init() {
    startTime = time.Now()
}

func waitForTerminationSignal() {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan
}
