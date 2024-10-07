# Theory about functions:

## Variadic Functions :
Variadic function are the functions that can take any number of arguments of a define type
Declaration :  func <func_name> (param-1 type, param-2 type,para-3 ...type)<return_type>
Example : func sumNumbers (numbers ...int) 
    - Now this function can take any number of arguments of type "int"
Example : func sumNumber (str string, number ...int)
- Variadic Functoins : A function that can take n number of parameters of same type
### Example : 
    ```
    func variadic(v...int) {
        fmt.Println("Elements are", v)
    }
    here in above example v will be treated as slice and can take 0 to n number of arguments
    ```
- Variadic functions can also be called with a slice, using the ... operator to unpack the slice into individual arguments
### Example : 
    ```
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(sum(nums...)) // Outputs: 15
    Its main examples can be Println() function of golang
    ```


## Anonymous functions : 
The functions where the function name is not given and are called just after they are defined.
Anonymous functions, also known as function literals or lambda expressions, are functions defined without a name. In Go, you can create anonymous functions inline, often as arguments to other functions or assigned to variables. They are particularly useful in scenarios where you need to define a small, one-off function without the need for a separate named function declaration.
### Example : 
```
func() {
    fmt.Println("This is anonymous function")
}() <- Here only we have called this function
```
Another example of Anonymous function : 
```
func main() {
// Defining an anonymous function and assigning it to a variable
add := func(x, y int) int {
    return x + y
}
// Calling the anonymous function via the variable
result := add(3, 4)
fmt.Println(result) // Outputs: 7
}
```
## Closure functions 
Closure functions, also known simply as closures, are functions that capture and retain the environment in which they are defined. This means they can access variables that are defined outside of their own scope. In Go, closures are implemented as anonymous functions.
    
### Example :
```
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
```
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
## Methods : 
In Go, a method is a function that is associated with a specific type. It allows you to define behaviors and actions that can be performed by values of that type. Methods are similar to functions but have a special receiver parameter that binds the method to the type.
This allows the method to access the receiver's fields and other methods, making it similar to object-oriented programming where methods belong to classes or objects.
```
-Syntax : 
func (receiver Type) methodName(parameters) returnType {
}
```

Receiver : The receiver is a parameter in the method declaration that specifies the type to which the method is bound. It associates the method with that type. The receiver can be of any named type (including struct types, pointer types, and even user-defined types). You can access the receiver's fields and call other methods on it inside the method's implementation.
- Methods can be defined on any type, including built-in types like int, string, and bool
- Methods can be defined on pointer types, which allows you to modify the receiver's fields inside the method
- Methods can be defined on user-defined types, which allows you to define custom behaviors and actions for

Difference between functions and methods : 
- Functions are not associated with any type, while methods are associated with a specific type.
- Methods can access the receiver's fields and call other methods on it, while functions cannot.
- Methods can be called on values of the type to which they are bound, while functions can be called on any value.
- Same methods can have different receivers and still works but same function defined twice with different parameters will give function redeclared error

There are two types of receivers : 
- value receivers that work on the copy values only
```
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

```
- Pointer receivers which work on original values
```
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```
## Interfaces and Methods:
Methods are also fundamental to implementing interfaces in Go. An interface specifies a set of method signatures, and a type implements an interface by providing definitions for all the methods in the interface's method set.
```
package main

import (
    "fmt"
)

// Define an interface
type Shape interface {
    Area() float64
}

// Define a struct type
type Rectangle struct {
    Width, Height float64
}

// Implement the interface by defining the Area method
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    var s Shape

    //Initialising struct
    r := Rectangle{Width: 5, Height: 10}

    // Assign a Rectangle to the Shape interface
    s = r

    // Call the Area method through the interface
    fmt.Println("Area:", s.Area()) // Output: Area: 50
}
```
Note : Same methods can have different receivers 
- Type: The type that the method is bound to. It can be a built-in type (e.g., int, string) or a user-defined type (e.g., a struct that you've defined).

Inferfaces : 
In Go, an interface is a collection of "method" signatures. It defines a set of methods that a type must implement to be considered as implementing that interface. In other words, an interface specifies a contract that a type must fulfill to be used in a particular context.
Interfaces enable polymorphism in Go, allowing you to write more flexible and reusable code. They help decouple the implementation details from the usage, making it easier to swap different types that satisfy the same interface without changing the code that uses them.

## Structs
Structs are a fundamental data type in Go. They are used to group related data together into a
single entity. Structs are similar to classes in other programming languages, but they are
more lightweight and do not have any built-in inheritance or encapsulation mechanisms.
There can be nested structs
### Example : 
```
type user struct{
    name string
    age int
    email string
    address interface{} //can take any value(data type)
}
```
They are initialised in following ways
```
var <name of variable> <name of struct>
var u user

Or we can directly assign values to them
u := user{"John", 25, "john@example.com", "123 Main St, AnyTown, USA"}
or another ways is
u := user{name: "John", age: 25, email: "john@example.com", address: "123 Main St, AnyTown, USA"}
```
In Go, anonymous structs are structs that are defined without a name and are typically used for short-lived data structures that are not reused elsewhere in the code. They can be useful for creating quick, temporary data structures, often within functions or when grouping multiple values together without the need to define a named struct type.
More examples in stucts.go
> Stucts in golang can have interfaces as well
```
type APIHandler struct {
    eventTracker     EventTracker
    affinityCalculator AffinityCalculator
    feedGenerator    FeedGenerator
}
```
here these are interfaces and have methods in them.
