# Context in Golang
In Go, context is a package (context) that provides a way to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes. It is commonly used for managing the lifecycle of a request, especially in server applications where you need to handle timeouts, cancellations, and other request-related metadata.

### 𝗪𝗵𝗲𝗻 𝘁𝗼 𝘂𝘀𝗲 𝗖𝗼𝗻𝘁𝗲𝘅𝘁:
- 𝗥𝗲𝗾𝘂𝗲𝘀𝘁 𝗧𝗶𝗺𝗲𝗼𝘂𝘁𝘀: To automatically cancel operations that take too long, like API calls or database queries.
- 𝗖𝗮𝗻𝗰𝗲𝗹𝗹𝗮𝘁𝗶𝗼𝗻 𝗣𝗿𝗼𝗽𝗮𝗴𝗮𝘁𝗶𝗼𝗻: To propagate cancellation signals across goroutines, ensuring all dependent operations stop if the parent operation is canceled.
- 𝗥𝗲𝗾𝘂𝗲𝘀𝘁-𝗦𝗰𝗼𝗽𝗲𝗱 𝗩𝗮𝗹𝘂𝗲𝘀: To pass request-scoped data, such as authentication tokens, between functions in a consistent manner.
- 𝗚𝗿𝗮𝗰𝗲𝗳𝘂𝗹 𝗦𝗵𝘂𝘁𝗱𝗼𝘄𝗻𝘀: To handle cleanup tasks during the shutdown of servers or other long-running processes.  
Example : 
```
func main(){
    fmt.Println("Process Started")

    //context with 2 seconds time-out
    ctx,cancel := context.WithTimeout(context.Backgroud(),2*time.Second)
    defer cancel() // ensures that the context is canceled when main() function exits, to avoid potential resource leaks.

    // simulate a process that take time
    ch:=make(chan string)
    go func(){
        time.Sleep(3 * time.Second)
        ch<-"Process completed"
    }()

    //wait for process or timeout

    select{
        case result:=<-ch:
            fmt.Println(result)
        case <-ctx.Done():
            fmt.Println("Operation timeout!")
    }

    fmt.Println("Process ended")
}

output : Process started
operation timed out!
Process ended
```