package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    // Password to be hashed
    password := "mySecretPassword123"

    // Generate a hashed representation of the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Original Password:", password)
    fmt.Println("Hashed Password:", string(hashedPassword))

    // Example of comparing the hashed password with a plaintext password
    // This typically happens during authentication
    inputPassword := "mySecretPassword123"
    err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(inputPassword))
    if err != nil {
        fmt.Println("Passwords do not match.")
        return
    }

    fmt.Println("Passwords match!")
}
