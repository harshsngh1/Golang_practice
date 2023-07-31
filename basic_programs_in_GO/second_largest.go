package main

import (
	"fmt"
	"math"
)

func SecondLargest(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	large := arr[0]
	secondLarge := math.MinInt

	for _, value := range arr {
		if value > large {
			secondLarge = large
			large = value
		} else if value > secondLarge && value != large {
			secondLarge = value
		}
	}
	return secondLarge
}

func main() {
	arr := []int{10, 5, 17, 8, 30}
	secondL := SecondLargest(arr)
	fmt.Println(secondL)
}
