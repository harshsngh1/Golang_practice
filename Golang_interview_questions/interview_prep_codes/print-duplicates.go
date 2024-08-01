package main

import "fmt"

// showDuplicates prints all duplicate elements only once
func showDuplicates(values []int) {
	// Create a map to store the count of each element
	valMap := make(map[int]int)

	// Count occurrences of each element in the input slice
	for _, value := range values {
		valMap[value]++
	}

	// Iterate over the map to find and print duplicates
	for key, count := range valMap {
		if count > 1 {
			fmt.Println(key)
		}
	}
}

func main() {
	// Example input slice
	slice := []int{1, 2, 3, 4, 5, 1, 2, 3, 3, 3}

	// Print duplicate elements
	showDuplicates(slice)
}
