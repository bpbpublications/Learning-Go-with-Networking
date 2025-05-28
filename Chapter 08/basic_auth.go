package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "encoding/base64"
)

func main() {
    // URL of the protected resource
    url := "https://api.example.com/data"
    
    // Username and password for authentication
    username := "your_username"
    password := "your_password"
    
    // Construct the HTTP request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating HTTP request:", err)
        os.Exit(1)
    }
    
    // Set Basic Authentication header
    req.SetBasicAuth(username, password)
    
    // Send the HTTP request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending HTTP request:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()
    
    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        os.Exit(1)
    }
    
    // Check if the request was successful (status code 200)
    if resp.StatusCode == 200 {
        fmt.Println("Request successful!")
        fmt.Println("Response:")
        fmt.Println(string(body))
    } else {
        fmt.Println("Request failed with status code:", resp.StatusCode)
    }
}
