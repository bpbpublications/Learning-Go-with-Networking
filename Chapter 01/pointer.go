package main

import "fmt"

func main() {
    // Declare variables
    num := 42
    str := "Hello"
    arr := [3]int{1, 2, 3}

    // Pointers to different types
    var intPtr *int
    var strPtr *string
    var arrPtr *[3]int

    // Assign addresses to pointers
    intPtr = &num
    strPtr = &str
    arrPtr = &arr

    // Print values and addresses
    fmt.Println("Value of num:", num)
    fmt.Println("Address of num:", &num)
    fmt.Println("Value stored in intPtr:", *intPtr)
    fmt.Println("Value of str:", str)
    fmt.Println("Address of str:", &str)
    fmt.Println("Value stored in strPtr:", *strPtr)
    fmt.Println("Value of arr:", arr)
    fmt.Println("Address of arr:", &arr)
    fmt.Println("Value stored in arrPtr:", *arrPtr)

    // Modify values indirectly
    *intPtr = 10
    *strPtr = "World"
    arrPtr[0] = 100

    // Print modified values
    fmt.Println("Modified value of num:", num)
    fmt.Println("Modified value of str:", str)
    fmt.Println("Modified value of arr:", arr)

    // Pointer to pointer
    var ptrToPtr **int
    ptrToPtr = &intPtr
    fmt.Println("Value stored in ptrToPtr:", **ptrToPtr)
}
