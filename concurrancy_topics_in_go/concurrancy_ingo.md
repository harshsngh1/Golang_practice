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
```
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
```
Key points about the select statement:

- It allows you to wait for multiple channel operations simultaneously.
- If multiple channels are ready, one of them is selected at random (uniform distribution).
- If no channel operation is ready and there's a default clause, the code inside the default block is executed. If there's no default   and no channel operation is ready, the select statement blocks until at least one of the channels becomes ready.

## Important Point about channels

- this code works fine
```
ch := make(chan int)
go multiply(ch)
ch<-10
```

- But this code breaks
```
ch := make(chan int)
ch<-10
go multiply(ch)
```

The reason for this is that we are using unbuffered channel which need a sender and receiver and here in 2nd code we are pushing data in unbuffered channel with no receiver (go routine), causing the main goroutine to block indefinitely.  As a result, the multiply goroutine never gets a chance to start, leading to a deadlock.
So to solve this we can use bufferend channel : 
```
ch := make(chan int, 1)  // Create a buffered channel with capacity 1
ch <- 10                 // This will not block because the channel has buffer space
go multiply(ch)
```
- Bidirectional channels : By default, channels in Go are bidirectional, meaning they can be used for both sending and receiving values.
```
func main() {
    // Create a bidirectional channel
    ch := make(chan int)

    // Goroutine to send a value into the channel
    go func() {
        ch <- 42
    }()

    // Receive the value from the channel
    value := <-ch
    fmt.Println("Received:", value)
}
```
- Unidirectional Channels : Unidirectional channels restrict the direction of data flow. A channel can be constrained to only send or only receive values. This is useful for making the intentions of your code clear and for avoiding accidental misuse of the channel.
-- Send-Only Channels : A send-only channel is declared by adding chan<- before the channel type. This restricts the channel to only sending values.
```
package main

import "fmt"

func sendValues(ch chan<- int) {
    ch <- 42
}

func main() {
    // Create a bidirectional channel
    ch := make(chan int)

    // Start a goroutine to send a value
    go sendValues(ch)

    // Receive the value from the channel
    value := <-ch
    fmt.Println("Received:", value)
}

In this example, the sendValues function takes a send-only channel (chan<- int). The function can send values into the channel but cannot receive from it.
```
Receive-Only Channels : A receive-only channel is declared by adding <-chan before the channel type. This restricts the channel to only receiving values.
```
package main

import "fmt"

func receiveValues(ch <-chan int) {
    value := <-ch
    fmt.Println("Received:", value)
}

func main() {
    // Create a bidirectional channel
    ch := make(chan int)

    // Start a goroutine to receive a value
    go receiveValues(ch)

    // Send a value into the channel
    ch <- 42
}
In this example, the receiveValues function takes a receive-only channel (<-chan int). The function can receive values from the channel but cannot send into it.
```
- There are 3 ways in which go routines can be controlled
-- wait groups
-- time.sleep()
-- synchronized channels : In Go, all channels are synchronized by default. This means operations on channels (sending and receiving) are blocking and provide thread-safe communication between goroutines.

- A good example
```
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for task := range tasks {
        fmt.Printf("Worker %d processing task %d\n", id, task)
        time.Sleep(time.Second) // Simulate work
    }
}

func main() {
    tasks := make(chan int, 10)
    var wg sync.WaitGroup

    // Start 3 workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, tasks, &wg)
    }

    // Send tasks
    for i := 1; i <= 5; i++ {
        tasks <- i
    }
    close(tasks) // Close the channel to signal no more tasks

    wg.Wait() // Wait for all workers to finish
}

```
- Mutex Locks : Mutex locks are a synchronization primitive used to prevent race conditions by ensuring that only one goroutine can access a shared resource at a time. In Go, the sync package provides the Mutex type, which is used to implement mutual exclusion.

Example : 

```
package main

import (
    "fmt"
    "sync"
)

var (
    count int
    mu    sync.Mutex
)

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    mu.Lock()   // Acquire the lock
    count++     // Critical section
    mu.Unlock() // Release the lock
}

func main() {
    var wg sync.WaitGroup

    // Start 10 goroutines
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go increment(&wg)
    }

    wg.Wait()
    fmt.Println("Final count:", count)
}

```
Note : Read-Write Mutex (sync.RWMutex): Allows multiple concurrent readers but only one writer at a time.