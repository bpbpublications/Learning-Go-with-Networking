package main

import "fmt"

func main() {
    // Creating an empty map
    ages := make(map[string]int)

    // Adding key-value pairs to the map
    ages["Alice"] = 28
    ages["Bob"] = 32
    ages["Charlie"] = 45

    // Accessing values from the map
    aliceAge := ages["Alice"]
    fmt.Println("Alice's age:", aliceAge)  // Output: Alice's age: 28

    // Modifying values in the map
    ages["Bob"] = 33
    fmt.Println("Bob's updated age:", ages["Bob"])  // Output: Bob's updated age: 33

    // Checking if a key exists in the map
    age, exists := ages["Charlie"]
    if exists {
        fmt.Println("Charlie's age:", age)  // Output: Charlie's age: 45
    } else {
        fmt.Println("Charlie's age not found")
    }

    // Deleting a key-value pair from the map
    delete(ages, "Alice")
    fmt.Println("Alice's age after deletion:", ages["Alice"])  // Output: Alice's age after deletion: 0

    // Iterating over the map
    for name, age := range ages {
        fmt.Printf("%s is %d years old\n", name, age)
    }
}
