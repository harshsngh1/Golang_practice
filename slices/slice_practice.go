package main

import "fmt"

func main() {
	slice := make([]int, 5, 8)
	slice_2 := []int{}

	// append method usage
	slice = append(slice, 10, 20, 30, 40)

	fmt.Println(slice)
	fmt.Println(slice_2)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
