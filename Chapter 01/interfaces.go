package main

import "fmt"

// Car interface
type Car interface {
    Start()
    Stop()
}

// Sedan struct implementing the Car interface
type Sedan struct {
    Brand string
}

// Start method implementation for Sedan
func (s Sedan) Start() {
    fmt.Println(s.Brand, "Sedan is starting...")
}

// Stop method implementation for Sedan
func (s Sedan) Stop() {
    fmt.Println(s.Brand, "Sedan is stopping...")
}

// SUV struct implementing the Car interface
type SUV struct {
    Brand string
}

// Start method implementation for SUV
func (s SUV) Start() {
    fmt.Println(s.Brand, "SUV is starting...")
}

// Stop method implementation for SUV
func (s SUV) Stop() {
    fmt.Println(s.Brand, "SUV is stopping...")
}

func main() {
    // Create Sedan and SUV instances
    sedan := Sedan{Brand: "Honda"}
    suv := SUV{Brand: "Toyota"}
    
    // Start and stop Sedan
    sedan.Start()
    sedan.Stop()
    
    // Start and stop SUV
    suv.Start()
    suv.Stop()
}
