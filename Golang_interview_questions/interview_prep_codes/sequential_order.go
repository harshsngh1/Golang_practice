package main

import (
	"fmt"
	"sync"
)

var num int32
var wg sync.WaitGroup
var mu sync.Mutex

// incrementNum increments num and prints the result
func incrementNum(goRoutineId int, ch chan struct{}, doneCh chan struct{}) {
	defer wg.Done()

	// Wait for the signal to start
	<-ch

	mu.Lock()
	num++
	fmt.Printf("goroutine %d incremented num to %d\n", goRoutineId, num)
	mu.Unlock()

	// Signal the next goroutine
	doneCh <- struct{}{}
}

func main() {
	maxVal := 100
	num = 0

	ch := make(chan struct{}, 1)
	doneCh := make(chan struct{})

	// Start the sequence
	ch <- struct{}{}

	// Launch goroutines
	for i := 1; i <= maxVal; i++ {
		wg.Add(1)
		go incrementNum(i, ch, doneCh)
		// Wait for the current goroutine to finish before allowing the next one to start
		<-doneCh
		if i < maxVal {
			ch <- struct{}{} // Signal the next goroutine to start
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the channels
	close(ch)
	close(doneCh)

	fmt.Printf("Value of num is %d\n", num)
}
