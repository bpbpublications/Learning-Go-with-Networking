package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "sync"
    "time"
)

func main() {
    urls := []string{
        "https://www.google.com",
        "https://www.google.com",
        "https://www.openai.com",
    }

    // Create an HTTP client with connection pooling
    client := &http.Client{
        Transport: &http.Transport{
            MaxIdleConns:        10,               // Maximum number of idle connections to keep
            IdleConnTimeout:     30 * time.Second, // Idle connection timeout
            DisableKeepAlives:   false,            // Enable keep-alive connections
            TLSHandshakeTimeout: 10 * time.Second, // TLS handshake timeout
        },
    }

    // Create a WaitGroup to track the completion of goroutines
    var wg sync.WaitGroup
    wg.Add(len(urls))

    // Make concurrent HTTP GET requests
    for _, url := range urls {
        go func(url string) {
            defer wg.Done()

            // Create an HTTP GET request
            req, err := http.NewRequest("GET", url, nil)
            if err != nil {
                fmt.Printf("Error creating request for %s: %s\n", url, err)
                return
            }

            // Send the request using the client
            resp, err := client.Do(req)
            if err != nil {
                fmt.Printf("Error sending request for %s: %s\n", url, err)
                return
            }
            defer resp.Body.Close()

            // Read and print the response body
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                fmt.Printf("Error reading response body for %s: %s\n", url, err)
                return
            }
            fmt.Printf("Response from %s: %s\n", url, string(body))
        }(url)
    }

    // Wait for all goroutines to finish
    wg.Wait()
}