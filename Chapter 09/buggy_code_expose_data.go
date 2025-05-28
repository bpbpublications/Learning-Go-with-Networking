package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Config struct represents the application configuration.
type Config struct {
	APIKey string `json:"api_key"`
	// Other configuration fields...
}

func loadConfig(filename string) (Config, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(fileContent, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func main() {
	configFile := "config.json"
	config, err := loadConfig(configFile)
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Simulate using the API key (e.g., making an API request)
	fmt.Println("API Key:", config.APIKey)
	// Actual application logic...
}
