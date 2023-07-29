package main

import "fmt"

func main() {
	mymap := make(map[string]int)
	fmt.Println(mymap)

	// creating a map and checking the value and flound concept

	codes := map[string]int{"en": 1, "fr": 2, "hi": 3}
	value, found := codes["fr"]
	fmt.Println(value, "is found?", found)

	// adding a value to current map

	codes["tl"] = 4

	fmt.Println(codes)

	// deleting a key in maps
	delete(codes, "tl")

	fmt.Println(codes)

	// itereate over maps
	for key, value := range codes {
		fmt.Println(key, "=>", value)
	}

	// truncate a map using new initialization

	codes = make(map[string]int)
	fmt.Println(codes)

	// map declared using make

	code2 := make(map[string]int)
	code2["en"] = 1

	fmt.Println(code2)
}
