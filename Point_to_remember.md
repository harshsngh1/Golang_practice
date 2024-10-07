# Points to remember : 
- golang is procedural language
- Staticly typed language, generates binary and then that binanry is executed
- "len" function can be used to find length of dervied data types like maps,slices,array
- Slices and maps are always passed by reference in a function by default
- channels are also always passed by reference in a function just like slices and maps
- "defer" keyword : delays the execution of a function or statement until the nearby function returns
- Go now supports "generic" from go 1.18 update
- Range loop is used on slice in golang, which has two values i,j = range <name of slice/array/strings>
    Here i will have the index and j will have the value
- For loop range can used on channels as well.
    ### Example :
    ```
    chnl := make(chan int)
    go func() {
        chnl <- 10
        chnl <- 100
        chnl <- 1000
        chnl <- 10000
        close(chnl)
    } ()
    for i:=range chnl{
        fmt.Println(i)
    }
    This is how we can do it.
    ```
- For loop/range loop can be used on objects as well in golang.
    ### Example : 
    ```
    mmap := map[int]string{
        0:"a",
        1:"b",
        2:"c",
    }
    for key,value: range mmap {
        fmt.Println(key,vale)
    }
    ```
- Init Function : The init() function is called automatically by the Go runtime before the main() function in the same package, or in the order of package import if there are multiple init() functions in different packages.
It neither takes any arguments and not have any return type.
They are used to initialise DB connections or establish some of the objects before initiation
- Main function is also a go routine
- Go routines are executed concurrently
- Go routines are executed in a thread pool

### Panic and recover Keyword : 
In Go, panic and recover are mechanisms used for handling unexpected conditions and errors in a program. They provide a way to deal with situations that are not typically managed through regular error handling.
- Panic : It is used to stop the execution of the program and print the error message
- When a function calls panic, the function's execution stops immediately, and any deferred functions are executed. The program then unwinds the stack, running any deferred functions until it reaches the top of the goroutine's stack. At this point, the program exits with a non-zero exit status.
- panic is typically used to indicate an unrecoverable error, such as a programming mistake (e.g., accessing an out-of-bounds index in a slice) or an unexpected state that should cause the program to terminate.
- There are two types of panic : custom and system
- Custom panic : It is used to create a custom panic message and exit the program
- System panic : It is used to handle system panics, which are caused by the Go runtime


- Recover : It is used to recover the program from panic and continue the execution of the program
- Recover is used inside a defer function
- Recover can only be called from within a deferred function. If it is called outside of a deferred
function, it will panic with the value nil.
- Recover can only be called once per panic. If it is called more than once, it will
panic with the value nil.
- When called within a deferred function, recover captures the value passed to panic and stops the panicking sequence, allowing the program to continue running.
- recover is used to handle panics gracefully, enabling the program to recover from unexpected conditions without crashing.
### Example
```
func safeDivide(a, b int) int {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

func main() {
    fmt.Println("Result:", safeDivide(4, 2)) // Output: Result: 2
    fmt.Println("Result:", safeDivide(4, 0)) // Output: Recovered from panic: division by zero
                                            //         Result: 0
}
```
### File Handling concepts
> Buffer Reader and Normal Reader:
- Buffered reader ek intermediate buffer ka use karta hai jo file se data ko chunks me load karta hai aur phir in-memory buffer se read karta hai, Buffered reader performance ko improve karta hai kyunki wo multiple reads ko combine karta hai ek hi file system access me. Data pehle buffer me load hota hai aur phir buffer se read hota hai, jo fast hota hai.
```
Buffered reader me aap bufio package ka NewReader ya NewScanner use karte hain.

package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
)

func main() {
    // File ko open karna
    file, err := os.Open("logfile.txt")
    if err != nil {
        log.Fatalf("Failed to open file: %s", err)
    }
    defer file.Close()

    // Buffered reader ka use karna
    scanner := bufio.NewScanner(file)
    
    lineNumber := 0
    for scanner.Scan() {
        lineNumber++
        // Har line ko read karna
        line := scanner.Text()
        // Process the line (aap apne hisab se process kar sakte hain)
        fmt.Println(line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading file: %s", err)
    }

    fmt.Printf("Total lines read: %d\n", lineNumber)
}
```
- Normal reader directly file se data ko read karta hai bina kisi buffer ke. Jab aap normal reader use karte hain, to har baar jab bhi aap read operation karte hain, wo file system se ek access request karta hai.

- GOMAXPROCS: Controls the number of OS threads used for executing goroutines.



Chatgpt learning : https://chat.openai.com/share/ce2a9c9b-b448-4684-ab11-a1e96d9b3549
