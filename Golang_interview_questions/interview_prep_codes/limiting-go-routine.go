package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	maxGoroutines := 100
	guard := make(chan struct{}, maxGoroutines) //buffered channel

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		guard <- struct{}{} // Block if there are already 100 goroutines running
		// Empty Struct: struct{}{} is an empty struct, which is used here because it takes up zero bytes of memory.
		// It's a way to signal without carrying any data.
		// When a goroutine wants to start, it sends an empty struct to the guard channel.
		// If the number of goroutines running is less than maxGoroutines,
		// the send operation (guard <- struct{}{}) will succeed immediately.
		// If there are already maxGoroutines goroutines running, the send operation will block until another
		// goroutine completes and releases a spot.

		go func(i int) {
			defer wg.Done()
			defer func() { <-guard }() // Release a spot in the semaphore
			time.Sleep(1 * time.Second)
			fmt.Println("Completed task", i)
		}(i)
	}

	wg.Wait()
}
