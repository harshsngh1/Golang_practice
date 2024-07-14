package interview_prep

import (
	"fmt"
	"sync"
)

func printOdd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd:", i)
		ch <- i // Send odd number to the channel
		<-ch    // Wait for acknowledgment from the even goroutine
	}
}

func printEven(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-ch // Wait for odd number from the channel
		fmt.Println("Even:", i)
		ch <- i // Send even number back to the channel
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go printOdd(ch, &wg)
	go printEven(ch, &wg)
	wg.Wait()
	close(ch)
}
