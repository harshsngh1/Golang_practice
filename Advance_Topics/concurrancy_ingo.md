# Notes :

- to lauch a go routine go have a key work "go".
  Example : go <func name>. go calculate()

## Go Schedular
- Golang scheduler is a cooperative scheduler.
- Cooperative scheduling is a style of scheduling in which the OS never interrupts a running
  process to initiate a context switch from one process to another. 
- Processes must voluntarily yield control periodically or when logically blocked on a
  resource.
- Of course, there are some specific check points where go-routine can yield its execution
  to other go-routine. These are called context switches
- It is upto schedular to decide to do weather context switching is requried or not.
- channels by design prevent race condition from hapenning

## Wait Groups
- The problem with go-routines was the main go-routine terminating before the go-routines completed or even began their execution.
- To wait for multiple go-routines to finish, we can use a wait group
- A wait group is a synchronization primitive that allows multiple go-routines to wait for each other
- In Golang (Go), a wait group is a synchronization mechanism used to control the execution flow of goroutines (concurrently executing functions). It allows you to wait for a group of goroutines to finish their execution before continuing with the rest of the program. Wait groups are particularly useful when you have multiple goroutines performing independent tasks, and you want to ensure that all of them complete before proceeding.
- The sync package in Go provides the WaitGroup type for managing wait groups

### Process for using wait groups
-- Create a WaitGroup: To use a wait group, you first need to create an instance of sync.WaitGroup.

-- Add goroutines to the WaitGroup: Before launching each goroutine that you want to wait for, you call the Add() method on the WaitGroup to increase the internal counter. The counter represents the number of goroutines the WaitGroup is waiting for.

-- Goroutines call Done(): Inside each goroutine, when the task is completed, you call the Done() method on the WaitGroup. This decrements the internal counter, signaling that the goroutine has finished its work.

-- Wait for all goroutines to finish: Finally, you use the Wait() method on the WaitGroup, which blocks the execution of the calling goroutine until the internal counter becomes zero. Once all goroutines have called Done(), the WaitGroup's counter becomes zero, and the Wait() method unblocks, allowing the program to continue executing.

## Channels


In Go, channels are a core feature used for communication and synchronization between goroutines (concurrently executing functions). They provide a way for goroutines to send and receive data to and from each other, facilitating safe and efficient communication in a concurrent program.
A channel is like a pipeline through which data can flow. Goroutines can send data into a channel and receive data from it. Channels ensure that data is transferred in a synchronized manner, allowing goroutines to safely communicate without race conditions.

Channels can be used to:
Share data between different goroutines in a thread-safe way.
Synchronize the execution of goroutines.
Implement patterns like producer-consumer, fan-out/fan-in, and more.

Syntax to create channel/unbuffered channel : ch := make(chan int) // Creates an unbuffered channel of integers

Unbuffered channels:
An unbuffered channel has no capacity to hold data. When data is sent through an unbuffered channel, the sender will block until there is a receiver ready to receive the data. Similarly, the receiver will block until there is a sender ready to send the data.

Buffered channels:
A buffered channel has a capacity that specifies the number of elements it can hold. Sending data to a buffered channel will only block if the channel is full, and receiving data will only block if the channel is empty.

syntax : ch := make(chan int, 10) // Buffered channel of integers with capacity 10

Channel send and receive:
To send data into a channel, you use the <- operator with the channel variable on the left-hand side. To receive data from a channel, you use the same operator with the channel variable on the right-hand side. For example:
ch <- data     // Send data into the channel
receivedData := <-ch // Receive data from the channel


## Select  Statements : 
In Go, a select statement is used to work with multiple channels, allowing you to wait for and perform operations on multiple channel communications simultaneously. It's a powerful construct that enables you to synchronize and interact with multiple goroutines efficiently.
The select statement looks somewhat like a switch statement but is specifically designed for dealing with channel operations

# Syntax
select {
case <-channel1:
    // Code to execute when data is received from channel1
case data := <-channel2:
    // Code to execute when data is received from channel2
case channel3 <- value:
    // Code to execute when data is sent to channel3
default:
    // Code to execute when no channel operation is ready
}

Key points about the select statement:

- It allows you to wait for multiple channel operations simultaneously.
- If multiple channels are ready, one of them is selected at random (uniform distribution).
- If no channel operation is ready and there's a default clause, the code inside the default block is executed. If there's no default   and no channel operation is ready, the select statement blocks until at least one of the channels becomes ready.