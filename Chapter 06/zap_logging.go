package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	// Create a logger configuration
	config := zap.NewProductionConfig()

	// Customize the logger configuration as needed
	config.OutputPaths = []string{"app.log"} // Output to a file
	config.ErrorOutputPaths = []string{"error.log"}

	// Create a logger instance
	logger, err := config.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync() // Flushes buffer and stops logging

	// Example application code
	for i := 0; i < 100000; i++ {
		// Perform some time-consuming operation
		result := timeConsumingOperation(i)

		// Log the result
		logger.Info("Iteration complete",
			zap.Int("iteration", i),
			zap.Int("result", result),
		)
	}
}

func timeConsumingOperation(i int) int {
	// Simulate a time-consuming operation
	time.Sleep(10 * time.Millisecond)

	// Return a result based on the input
	return i * 2
}