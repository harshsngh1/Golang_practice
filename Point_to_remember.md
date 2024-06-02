# Points to remember : 
- golang is procedural language
- Staticly typed language, generates binary and then that binanry is executed
- "len" function can be used to find length of dervied data types like maps,slices,array
- Slices and maps are always passed by reference in a function by default
- channels are also always passed by reference in a function just like slices and maps
- "defer" keyword : delays the execution of a function or statement until the nearby function returns
- String defined using backticks [``] are called raw literals and are different from the ones defined usinging double quotes, as in strings with back ticks the escape charecters like \n or \t are not considered and are treated as part of string only
    ### Example : s := `sdasdsadsadasd \n adsaasasas`
    here above s will have the whole word with it.
- Strings are immutable in golang, they cannot be changed
- Range loop is used on slice in golang, which has two values i,j = range <name of slice/array/strings>
    Here i will have the index and j will have the value
- For loop range can used on channels as well.
    ### Example :
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
- For loop/range loop can be used on objects as well in golnag.
    ### Example : 
    mmap := map[int]string{
        0:"a",
        1:"b",
        2:"c",
    }
    for key,value: range mmap {
        fmt.Println(key,vale)
    }
- Variadic Functoins : A function that can take n number of parameters of same type
    ### Example : 
    func variadic(v...int) {
        fmt.Println("Elements are", v)
    }
    here in above example v will be treated as slice and can take 0 to n number of arguments
- Variadic functions can also be called with a slice, using the ... operator to unpack the slice into individual arguments
    ### Example : 
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(sum(nums...)) // Outputs: 15
    Its main examples can be Println() function of golang

- Anonymous functions : The functions where the function name is not given and are called just after they are defined.
    Anonymous functions, also known as function literals or lambda expressions, are functions defined without a name. In Go, you can create anonymous functions inline, often as arguments to other functions or assigned to variables. They are particularly useful in scenarios where you need to define a small, one-off function without the need for a separate named function declaration.
    ### Example : 
    func() {
        fmt.Println("This is anonymous function")
    }() <- Here only we have called this function

    ### Another example of Anonymous function : 
    func main() {
    // Defining an anonymous function and assigning it to a variable
    add := func(x, y int) int {
        return x + y
    }

    // Calling the anonymous function via the variable
    result := add(3, 4)
    fmt.Println(result) // Outputs: 7
    }

- Init Function : The init() function is called automatically by the Go runtime before the main() function in the same package, or in the order of package import if there are multiple init() functions in different packages.
It neither takes any arguments and not have any return type.

- Closure functions : Closure functions, also known simply as closures, are functions that capture and retain the environment in which they are defined. This means they can access variables that are defined outside of their own scope. In Go, closures are implemented as anonymous functions.
    ### Example : 
    func main() {
    x := 10

    // Closure function that captures variable x
    increment := func() {
        x++
        fmt.Println(x)
    }

    increment() // Outputs: 11
    increment() // Outputs: 12
    }
    Closures are often used when you want to create functions that exhibit behavior based on some state that persists across multiple function calls. They are particularly useful in scenarios such as event handlers, callback functions, or when creating iterators.

    ### Another Example : 
```
    func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
        }
    }

    func main() {
        next := counter()
        fmt.Println(next()) // Outputs: 1
        fmt.Println(next()) // Outputs: 2
        fmt.Println(next()) // Outputs: 3
    }
```
- Methods and functions are different in Golang, methods have receivers whereas functions dont have receivers
Example of method : 
```
func (v values) add(a,b int) int {
    return a + b + v.first + v.second
}

here values is a struct having first and second and  (v values) is receiver and add is name of method
The receiver v is an instance of the values type, meaning the method can access the fields and other methods of values.
```






Chatgpt learning : https://chat.openai.com/share/ce2a9c9b-b448-4684-ab11-a1e96d9b3549
