package main

import (
	"fmt"
	"sync"
)

// getDigit returns a fixed digit
func getDigit() int {
	return 5
}

func main() {
	// Number of times to call getDigit concurrently
	const numCalls = 5

	// Channel to collect results from each goroutine
	results := make(chan int, numCalls)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launching goroutines
	// whenever there is case that we have to call a function which has a return type and we cannot modify it
	// we can use this approach of anonymous function to make this happen
	// it acts as a wrapper function on that function
	for i := 0; i < numCalls; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := getDigit()
			results <- result
		}()
	}

	// Waiting for all goroutines to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Summing the results
	sum := 0
	for result := range results {
		sum += result
	}

	fmt.Println("Sum of results:", sum)
}
