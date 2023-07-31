package main

import "fmt"

// Define an interface named Shape
type Shape interface {
	Area() float64
}

// Define a custom type Rectangle that satisfies the Shape interface
type Rectangle struct {
	width  float64
	height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Define a custom type Circle that satisfies the Shape interface
type Circle struct {
	radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	// Create instances of Rectangle and Circle
	rect := Rectangle{width: 5.0, height: 3.0}
	circle := Circle{radius: 2.5}

	// Use the Area method via the Shape interface
	fmt.Println("Area of the rectangle:", getArea(rect))
	fmt.Println("Area of the circle:", getArea(circle))
}

// Function that takes a Shape interface and calculates the area
func getArea(s Shape) float64 {
	return s.Area()
}

//In this example, we define the Shape interface with a single method Area(). Then, we create two custom types Rectangle and Circle,
//both of which implement the Shape interface by providing their respective Area() methods.

//In the main function, we demonstrate polymorphism by calling the getArea function with both Rectangle and Circle instances.
//Since both types satisfy the Shape interface, we can pass them as arguments to getArea, and the appropriate Area() method will be called for each type, returning the correct area.
