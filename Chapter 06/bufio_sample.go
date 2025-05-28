package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func blockingIOOperation() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter something: ")
    text, _ := reader.ReadString('\n')
    fmt.Println("You entered:", text)
}

func nonBlockingIOOperation() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Non-blocking operation: I'm doing something else...")
    time.Sleep(2 * time.Second)
    if reader.Buffered() > 0 {
        text, _ := reader.ReadString('\n')
        fmt.Println("You entered:", text)
    } else {
        fmt.Println("No input received")
    }
}

func main() {
    fmt.Println("Start")

    fmt.Println("Blocking I/O operation:")
    blockingIOOperation()

    fmt.Println("Non-blocking I/O operation:")
    nonBlockingIOOperation()

    fmt.Println("End")
}