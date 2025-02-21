package main

import "fmt"

func main() {
	c := make(chan string)
	initString(c)

	for {
		resp, ok := <-c
		if !ok { // Channel is closed
			fmt.Println("Channel Closed", ok)
			break
		}
		fmt.Println("Channel Open", resp, ok)
	}
}

func initString(chnl chan string) {
	for v := 0; v < 3; v++ {
		chnl <- "go guruji"
	}
	close(chnl)
}

// The above code is a simple example of how to use the ok syntax to check if a channel is open or closed.
// The ok syntax is a way to check if a channel is open or closed.
// If the channel is open, the ok syntax will return true and the value will be sent to the channel.
// If the channel is closed, the ok syntax will return false and the value will not be sent to the channel.
// This is useful when we want to check weather channel is closed or not.

/* But this code is having some issues :
- it will give error : fatal error: all goroutines are asleep - deadlock!
- This happens because the main goroutine is waiting for data from the channel, but the channel is not being written to by any other goroutine.
- This is a classic deadlock situation in Go.
- To fix this, we need to add a goroutine that writes to the channel. Means that we need to add a goroutine that writes to the channel in the initString function.
- As in main func we have called initString func and initString func has a channel which is a unbuffered channel that is why it is waiting for the main func to receive the data from the channel.
- So to fix this we need to add a goroutine that writes to the channel in the initString func or use a buffered channel.
*/
