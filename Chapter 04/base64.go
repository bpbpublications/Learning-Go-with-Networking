package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, World!")

	// Encoding to base64
	encodedData := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded Data:", encodedData)

	// Decoding from base64
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Decoded Data:", string(decodedData))
}