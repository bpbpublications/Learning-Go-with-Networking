package main

import "fmt"

// Define a struct type for a car
type Car struct {
    Brand    string
    Model    string
    Year     int
    Engine   float64
    Features []string
    Options  map[string]bool
    Start    func()
    Stop     func()
}

func main() {
    // Create an instance of a car
    myCar := Car{
        Brand:  "Toyota",
        Model:  "Camry",
        Year:   2022,
        Engine: 2.5,
        Features: []string{"ABS", "Airbags", "GPS"},
        Options: map[string]bool{
            "Sunroof":   true,
            "AC":        true,
            "Bluetooth": true,
        },
        Start: func() {
            fmt.Println("Engine started!")
        },
        Stop: func() {
            fmt.Println("Engine stopped!")
        },
    }

    // Accessing
