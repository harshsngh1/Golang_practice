package main

import (
	"fmt"
)

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	// Initialize using composite literals -- way one
	john := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street: "123 Elm St",
			City:   "Somewhere",
			State:  "CA",
			Zip:    "90001",
		},
	}

	fmt.Printf("Person: %+v\n", john)
}

func main() {
	// Initialize using separate assignments -- way two
	var jane Person
	jane.Name = "Jane Smith"
	jane.Age = 28
	jane.Address.Street = "456 Oak St"
	jane.Address.City = "Elsewhere"
	jane.Address.State = "TX"
	jane.Address.Zip = "75001"

	fmt.Printf("Person: %+v\n", jane)
}

// Nesting Structs with Pointers:
// You can also use pointers to nested structs, which is useful for large structs or when you want to avoid copying the entire struct.

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Person struct {
	Name    string
	Age     int
	Address *Address
}

func main() {
	// Initialize using pointers
	homeAddress := &Address{
		Street: "789 Pine St",
		City:   "Anywhere",
		State:  "NY",
		Zip:    "10001",
	}

	bob := Person{
		Name:    "Bob Brown",
		Age:     35,
		Address: homeAddress,
	}

	fmt.Printf("Person: %+v\n", bob)
}

// Anonymous structs example
// they are initialised as soon as they are defined
package main

import "fmt"

func main() {
    // Define and initialize an anonymous struct
    person := struct {
        Name string
        Age  int
    }{// initialised here only
        Name: "Alice",
        Age:  30,
    }

    fmt.Printf("Person: %+v\n", person)
}
