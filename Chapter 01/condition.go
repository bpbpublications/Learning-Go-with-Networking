package main

import "fmt"

func main() {
    // Example 1: Basic if statement
    num := 10
    if num > 0 {
        fmt.Println("Number is positive")
    }

    // Example 2: if-else statement
    num = -5
    if num > 0 {
        fmt.Println("Number is positive")
    } else {
        fmt.Println("Number is negative")
    }

    // Example 3: if-else if-else statement
    num = 0
    if num > 0 {
        fmt.Println("Number is positive")
    } else if num < 0 {
        fmt.Println("Number is negative")
    } else {
        fmt.Println("Number is zero")
    }

    // Example 4: Nested if statements
    num = 10
    if num > 0 {
        if num%2 == 0 {
            fmt.Println("Number is positive and even")
        } else {
            fmt.Println("Number is positive and odd")
        }
    }

    // Example 5: Short-circuit evaluation
    num1 := 5
    num2 := 10
    if num1 > 0 && num2 > 0 {
        fmt.Println("Both numbers are positive")
    }

    // Example 6: Inline if statement
    isTrue := true
    if val := getValue(); val > 0 {
        fmt.Println("Value is positive")
    } else if isTrue {
        fmt.Println("Value is negative")
    }
}

// Function to return a value
func getValue() int {
    return -5
}
