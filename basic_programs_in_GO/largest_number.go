package main

import "fmt"

func Largestnum(arr []int) int {
	maxnum := arr[0]

	for _, value := range arr {
		if value >= maxnum {
			maxnum = value
		}
	}
	return maxnum
}

func main() {
	arr := []int{10, 5, 17, 8, 30}
	largest := Largestnum(arr)
	fmt.Println(largest)
}
