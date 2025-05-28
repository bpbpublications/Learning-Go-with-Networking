package main

import (
    "fmt"
    "geomentry"
)

func main() {
    // Using the exported function
    area := geomentry.CalculateCircleArea(2.5)
    fmt.Println("Circle area:", area)

    // Using the exported constant
    fmt.Println("Max attempts:", geomentry.MaxAttempts)
    fmt.Println("Error message:", geomentry.ErrorMessage)

    // Using the exported type
    rect := geomentry.Rectangle{Width: 4.5, Height: 2.5}
    area = rect.CalculateArea()
    fmt.Println("Rectangle area:", area)
}
