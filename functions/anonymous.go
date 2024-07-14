package main

import "fmt"

func main() {
	// anonymous function, since it is returning a value
	// we have assigned it to a variable
	x := func(l int, b int) int {
		return l * b
	}
	fmt.Printf("%T \n", x)
	fmt.Println(x(20, 30))

	// another example
	y := func(l int, b int) int {
		return l * b
	}(20, 30)
	fmt.Printf("%T \n", y)
	fmt.Println(y)

}

//another example

/*
package main

import "fmt"

func main() {
    // Initialize and assign an anonymous function to the variable "add"
    add := func(a, b int) int {
        return a + b
    }

    // Call the anonymous function through the "add" variable
    result := add(3, 5)
    fmt.Println(result) // Output: 8
}

func main() {
    println((func(x, y int) int {
        return x + y
    })(3, 4)) // Output: 7
}

Another Example

func() {
fmt.Println("anonymous function")
}()
called here only

*/
