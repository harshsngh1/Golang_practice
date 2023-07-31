package main

import "fmt"

// Define a struct type
type Rectangle struct {
	width  float64
	height float64
}

// Method to calculate the area of the Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	// Create a Rectangle instance
	rect := Rectangle{width: 5.0, height: 3.0}

	// Call the Area method on the Rectangle instance
	area := rect.Area()

	fmt.Println("Area of the Rectangle:", area) // Output: Area of the Rectangle: 15
}
