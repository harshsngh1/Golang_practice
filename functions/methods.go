package main

import (
	"fmt"
)

// Define a struct type
type Rectangle struct {
	Width, Height float64
}

// Define a method with Rectangle as the receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Define another method with Rectangle as the receiver
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {

	//here r is instance of struct and below we have used the methods with r as instance
	r := Rectangle{Width: 5, Height: 10}

	// Call the methods
	fmt.Println("Area:", r.Area())           // Output: Area: 50
	fmt.Println("Perimeter:", r.Perimeter()) // Output: Perimeter: 30
}
