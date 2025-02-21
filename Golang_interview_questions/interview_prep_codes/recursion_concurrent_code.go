/*
Question :
Write a program to launch a number of goroutines recursively and send data back to the main goroutine
where from main go func we launch a go routine and then that go routine will launch another go routine
and so on till a condition is satisfied or lets say there is int in main func having any integer value
maybe 4,5 so in these many go routine will be launch in the above manner and then only data needs
to be shared via channels only and whatever will be the last routine,
it should share the data to main go routine
*/

package main

import (
	"fmt"
)

// Goroutine function that launches the next goroutine recursively
func launchGoroutines(n int, ch chan string) {
	if n == 0 {
		// Base case: last goroutine, send data back to main
		ch <- "Data from the last goroutine"
		return
	}

	// Create a new channel for the next goroutine
	nextCh := make(chan string)

	// Launch the next goroutine
	go launchGoroutines(n-1, nextCh)

	// Wait for data from the next goroutine
	data := <-nextCh

	// Send received data to the previous goroutine
	ch <- data
}

func main() {
	// Number of goroutines to launch
	n := 4

	// Create a channel for the first goroutine
	ch := make(chan string)

	// Launch the first goroutine
	go launchGoroutines(n, ch)

	// Wait for the final data from the last goroutine
	result := <-ch

	fmt.Println("Final data received in main:", result)
}
