# OOPS with Goang

While Go (Golang) is not traditionally considered an object-oriented programming (OOP) language like Java or C++, it does support many OOP concepts through its unique approach. Here are some key OOP concepts and how they can be implemented in Go:

## Encapsulation
Encapsulation is achieved in Go through the use of structs and methods. Fields and methods can be exported (public) or unexported (private) based on their capitalization.

```
package main

import "fmt"

// Person struct with exported and unexported fields
type Person struct {
    Name string // Exported field
    age  int    // Unexported field
}

// NewPerson is a constructor function for Person
func NewPerson(name string, age int) *Person {
    return &Person{Name: name, age: age}
}

// GetAge is a method to access the unexported age field
func (p *Person) GetAge() int {
    return p.age
}

func main() {
    person := NewPerson("Alice", 30)
    fmt.Println(person.Name) // Accessing exported field
    fmt.Println(person.GetAge()) // Accessing unexported field via method
}
```

## Inheritance
Go does not have classical inheritance. Instead, it uses composition to achieve similar results. Composition involves embedding structs within other structs.

```
package main

import "fmt"

// Base struct
type Animal struct {
    Name string
}

// Method associated with Animal
func (a *Animal) Speak() {
    fmt.Println(a.Name, "makes a sound")
}

// Derived struct with embedded Animal
type Dog struct {
    Animal
    Breed string
}

// Method overriding
func (d *Dog) Speak() {
    fmt.Println(d.Name, "barks")
}

func main() {
    dog := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden Retriever"}
    dog.Speak() // Buddy barks
}
```

## Polymorphism
Polymorphism in Go is achieved through interfaces. An interface is a set of method signatures that a type must implement.

```
package main

import "fmt"

// Speaker interface
type Speaker interface {
    Speak()
}

// Animal struct
type Animal struct {
    Name string
}

// Animal implements Speaker interface
func (a Animal) Speak() {
    fmt.Println(a.Name, "makes a sound")
}

// Dog struct
type Dog struct {
    Animal
}

// Dog overrides Speak method
func (d Dog) Speak() {
    fmt.Println(d.Name, "barks")
}

func main() {
    var s Speaker

    s = Animal{Name: "Generic Animal"}
    s.Speak() // Generic Animal makes a sound

    s = Dog{Animal: Animal{Name: "Buddy"}}
    s.Speak() // Buddy barks
}
```

## Abstraction
Abstraction is achieved through interfaces and by hiding implementation details in unexported fields and methods.

```
package main

import "fmt"

// Shape interface
type Shape interface {
    Area() float64
}

// Circle struct
type Circle struct {
    radius float64
}

// Circle implements Shape interface
func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

// Rectangle struct
type Rectangle struct {
    width, height float64
}

// Rectangle implements Shape interface
func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func main() {
    var s Shape

    s = Circle{radius: 5}
    fmt.Println("Area of Circle:", s.Area()) // Area of Circle: 78.5

    s = Rectangle{width: 4, height: 6}
    fmt.Println("Area of Rectangle:", s.Area()) // Area of Rectangle: 24
}
```
By using structs, methods, interfaces, and composition, Go allows you to implement OOP principles in a way that is idiomatic to the language.