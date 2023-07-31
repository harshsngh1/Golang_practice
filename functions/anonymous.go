package main

import "fmt"

func main() {
	// anonymous function
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
