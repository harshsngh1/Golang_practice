## Interview Questions
- Can you explain how Go's concurrency model works?
Go's concurrency model is built around Goroutines and channels. Goroutines are lightweight threads managed by the Go runtime, allowing concurrent execution of functions. Channels facilitate communication and synchronization between Goroutines. This model promotes a "Do not communicate by sharing memory; instead, share memory by communicating" philosophy, which helps avoid common concurrency issues like race conditions and deadlocks.

- How would you handle error handling in Go?
In Go, error handling is explicit and encourages the use of the error type. Functions that may return an error typically return it as the last value. Developers then check the returned error value and handle it accordingly. This approach ensures that errors are not ignored and promotes robust error handling throughout the codebase.

- Can you compare interfaces and structs in Go?
Interfaces define behavior, while structs define data structures. Interfaces are implicitly satisfied, meaning types in Go automatically satisfy an interface if they implement its methods. Structs, on the other hand, are composite types that group together variables of different types under a single name. Interfaces enable polymorphism, allowing different types to be treated uniformly based on their behavior, while structs are used for organizing data.

- Explain defer, panic, and recover in Go.
-- defer: Defers the execution of a function until the surrounding function returns. It's commonly used for cleanup actions like closing files or releasing resources.
-- panic: Indicates a runtime error and stops the normal flow of execution. It's often used to signal unrecoverable errors.
-- recover: Used to regain control of a panicking goroutine and resume normal execution. It's typically used in deferred functions to handle panics gracefully and allow the program to continue running.

- What are Goroutines and how do they differ from threads?
Goroutines are lightweight, managed by the Go runtime, and have a smaller stack size compared to threads. They're more efficient in terms of memory usage and can be multiplexed onto a smaller number of operating system threads. Goroutines are designed to be used in large numbers, allowing highly concurrent programs without the overhead associated with traditional threads.