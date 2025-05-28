package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("URL: %s, Length: %d\n", url, len(body))
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
	}

	for _, url := range urls {
		go fetchURL(url) // Concurrently download each URL
	}

	// Wait for all goroutines to finish
	fmt.Println("Waiting for downloads to complete...")
	fmt.Scanln()
	fmt.Println("Done")
}
