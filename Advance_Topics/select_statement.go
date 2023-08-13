package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 42
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "hello"
	}()

	select {
	case val := <-ch1:
		fmt.Println("Received from ch1:", val)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	case <-time.After(4 * time.Second):
		fmt.Println("Timeout!")
	}
}

/*
In this example, two goroutines send data into two different channels after a certain delay.
The select statement waits for the first channel to receive data, the second channel to receive data, or a timeout to occur.
Whichever operation is ready first will be selected, and the corresponding code block will execute.

*/
