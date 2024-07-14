// The select keyword in Go is used to handle multiple channel operations.
// It allows a goroutine to wait on multiple communication operations, proceeding with the first one that is ready.
// This is particularly useful for managing multiple channels concurrently and implementing patterns such as timeouts,
// multiplexing, and coordinating between different goroutines.
// How select Works:
// - Waiting on Multiple Channels: The select statement blocks until one of its cases can proceed, then it executes that case.
// - Random Selection: If multiple cases are ready, one is chosen at random to be executed.
// - Default Case: If a default case is provided, it executes immediately if no other case is ready.

// Syntax :

// select {
// case <-channel1:
//     // Code to execute when channel1 is ready
// case msg := <-channel2:
//     // Code to execute when channel2 is ready and msg received
// case channel3 <- value:
//     // Code to execute when value can be sent to channel3
// default:
//     // Code to execute if no channel is ready
// }

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

// Conclusion:
// The select statement is a powerful tool in Go for managing multiple channel operations,
// allowing a goroutine to wait on multiple channels and proceed with the first one that is ready.
// It helps in building responsive and concurrent applications by providing a simple yet flexible way to handle multiple communication events.
