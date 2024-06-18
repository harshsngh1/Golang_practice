package main

import (
	"fmt"
)

func printEven(ch chan int) {
	for i := 2; i <= 20; i += 2 {
		<-ch // Wait for the signal to print even number
		fmt.Println("even:", i)
		ch <- 1 // Signal that even number has been printed
	}
}

func printOdd(ch chan int) {
	for i := 1; i <= 19; i += 2 {
		fmt.Println("odd:", i)
		ch <- 1 // Signal that odd number has been printed
		<-ch    // Wait for the signal to print next odd number
	}
}

func main() {
	ch := make(chan int)
	go printEven(ch)
	go printOdd(ch)

	// Start the printing process by sending initial signal
	ch <- 1

	// Allow some time for goroutines to complete
	select {}
}
