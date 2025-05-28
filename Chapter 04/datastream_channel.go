 package main
  
  import (
  	"fmt"
  	"time"
  )
  
  // Concurrent processing using goroutines
  func processData(input <-chan int, output chan<- int) {
  	for num := range input {
  		// Perform some processing on the input data
  		result := num * 2
  
  		// Simulate some computational delay
  		time.Sleep(time.Millisecond * 500)
  
  		// Send the processed data to the output channel
  		output <- result
  	}
  }
  
  // Fan-out pattern to distribute data to multiple processing goroutines
  func fanOut(input <-chan int, numWorkers int) []chan int {
  	outputs := make([]chan int, numWorkers)
  	for i := 0; i < numWorkers; i++ {
  		outputs[i] = make(chan int)
  		go processData(input, outputs[i])
  	}
  	return outputs
  }
  
  // Fan-in pattern to merge results from multiple processing goroutines
  func fanIn(inputs []chan int, output chan<- int) {
  	for _, input := range inputs {
  		go func(input chan int) {
  			for num := range input {
  				// Send the processed data to the output channel
  				output <- num
  			}
  		}(input)
  	}
  }
  
  func main() {
  	// Create input and output channels
  	input := make(chan int)
  	output := make(chan int)
  
  	// Buffer the output channel for efficient streaming
  	bufferSize := 10
  	bufferedOutput := make(chan int, bufferSize)
  
  	// Step 1: Concurrent processing
  	numWorkers := 3
  	inputs := fanOut(input, numWorkers)
  
  	// Step 2: Fan-in pattern
  	go fanIn(inputs, output)
  
  	// Step 3: Apply backpressure with a rate limiter
  	limiter := time.Tick(time.Second)
  	go func() {
  		for num := range output {
  			<-limiter // Wait for rate-limited interval
  			bufferedOutput <- num
  		}
  	}()
  
  	// Step 4: Stream composition and processing
  	go func() {
      for num := range bufferedOutput {
       // Apply additional processing or transformations
       result := num + 10
    
       // Print the final processed data
       fmt.Println(result)
      }
     }()
    
     // Step 5: Send data through the input channel
     for i := 1; i <= 20; i++ {
      input <- i
     }
    
     // Close the input channel to signal the end of data
     close(input)
    
     // Wait for the processing to complete