package main

import (
    "bytes"
    "fmt"
    "net/http"
)

func main() {
    // Replace with your ZAP proxy address and port
    zapURL := "http://localhost:8080"

    // Example: Spider a target URL
    targetURL := "http://example.com"
    spiderURL := fmt.Sprintf("%s/JSON/spider/action/scan/?url=%s", zapURL, targetURL)

    resp, err := http.Post(spiderURL, "application/json", bytes.NewBuffer([]byte{}))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    // Handle the response as needed

    // Example: Retrieve the scan results
    // Replace SCAN_ID with the actual scan ID obtained from the Spider response
    scanID := "SCAN_ID"
    scanStatusURL := fmt.Sprintf("%s/JSON/spider/view/status/?scanId=%s", zapURL, scanID)

    resp, err = http.Get(scanStatusURL)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    // Handle the response as needed
}
