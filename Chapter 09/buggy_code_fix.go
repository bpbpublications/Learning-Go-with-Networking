package main
import (
	"fmt"
	"log"
	"os"
)
// Config struct represents the application configuration.
type Config struct {
	APIKey string
	// Other configuration fields...
}
func loadConfig() Config {
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatal("API_KEY environment variable is not set")
	}

	return Config{
		APIKey: apiKey,
		// Initialize other configuration fields...
	}
}
func main() {
	config := loadConfig()
	// Simulate using the API key (e.g., making an API request)
	fmt.Println("API Key:", config.APIKey)
	// Actual application logic...
}
