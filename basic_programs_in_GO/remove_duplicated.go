package main

import "fmt"

func removeDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// Initialize the last unique element position (i) to 0
	i := 0

	for j := 1; j < len(nums); j++ {
		// If the current element is different from the last unique element,
		// copy it to the next position after the last unique element
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	// Return the array up to (i+1), as elements after that are duplicates
	return nums[:i+1]
}

func main() {
	arr := []int{1, 1, 2, 3, 3, 4, 4, 4, 5, 6, 6}
	fmt.Println("Original Array:", arr)

	arr = removeDuplicates(arr)
	fmt.Println("Array with Duplicates Removed:", arr)
}
