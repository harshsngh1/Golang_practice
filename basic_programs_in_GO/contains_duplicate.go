package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// Create a hash set to store encountered elements
	seen := make(map[int]bool)

	for _, num := range nums {
		// If the element is already in the hash set, return true (duplicate found)
		if seen[num] {
			return true
		}

		// Add the element to the hash set
		seen[num] = true
	}

	// No duplicates found
	return false
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Array 1 contains duplicates:", containsDuplicate(arr1))

	arr2 := []int{1, 2, 3, 2, 4}
	fmt.Println("Array 2 contains duplicates:", containsDuplicate(arr2))
}

// leetcode version
/*
func containsDuplicate(nums []int) bool {
	nums_map := map[int]int{}
	for _, n := range nums {
		if _, ok := nums_map[n]; !ok {
			nums_map[n] = 1
		} else {
			return true
		}
	}
	return false

}

*/
