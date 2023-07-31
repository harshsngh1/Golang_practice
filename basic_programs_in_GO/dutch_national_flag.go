package main

import "fmt"

func sortColors(nums []int) {
	// Initialize three pointers to keep track of the positions
	// where 0s, 1s, and 2s should be placed.
	p0 := 0             // Pointer for 0s (leftmost)
	p2 := len(nums) - 1 // Pointer for 2s (rightmost)
	current := 0        // Current position for iteration

	// Loop through the array and move 0s to the left and 2s to the right,
	// leaving 1s in the middle.
	for current <= p2 {
		if nums[current] == 0 {
			// Swap with p0 and move both pointers forward.
			nums[current], nums[p0] = nums[p0], nums[current]
			current++
			p0++
		} else if nums[current] == 2 {
			// Swap with p2 and move p2 pointer backward.
			nums[current], nums[p2] = nums[p2], nums[current]
			p2--
		} else {
			// Skip 1s and move the current pointer forward.
			current++
		}
	}
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	fmt.Println("Original Array:", arr)

	sortColors(arr)
	fmt.Println("Sorted Array:", arr)
}
