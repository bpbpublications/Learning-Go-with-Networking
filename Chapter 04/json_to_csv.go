package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    data := [][]string{
        {"Name", "Age"},
        {"John", "25"},
        {"Alice", "30"},
        {"Bob", "40"},
    }

    file, err := os.Create("output.csv")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    writer.WriteAll(data)
    writer.Flush()

    if err := writer.Error(); err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("CSV file created successfully")
}