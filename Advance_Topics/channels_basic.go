package main

import (
	"fmt"
	"time"
)

// chan<- means sending data into channel
// <-chan means sending data out of channel
/*
Goroutines communicate through channels by using the <- operator.
The arrow points to the direction of data flow: <-ch means receiving from the channel ch
and ch <- data means sending data into the channel ch.
*/

// here we are creating an input channel where this channel "out" will receive 1 to 5 int in it and once data is sent, we close it.
func counter(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
		time.Sleep(time.Second)
	}
	close(out)
}

// here we have declared output channel where channel "in" will output the data received to it.
func printer(in <-chan int) {
	for num := range in {
		fmt.Printf("%d ", num)
	}
}

func main() {
	channel := make(chan int)
	go counter(channel) // lauched using go routine
	printer(channel)
}
