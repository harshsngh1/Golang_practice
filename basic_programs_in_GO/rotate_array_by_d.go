package main

import "fmt"

func rotateArray(arr []int, d int) []int {
	n := len(arr)
	d = d % n

	reverseArray(arr, 0, d-1)
	reverseArray(arr, d, n-1)
	reverseArray(arr, 0, n-1)

	return arr
}

func reverseArray(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	d := 3

	rotatedArray := rotateArray(arr, d)
	fmt.Println("Rotated Array:", rotatedArray)
}

//also in brute force approch to fill the elements back in main array you can start the itiration from n-d and continue
// n-d,n-d+1,n-d+2
// for shifting the existing element you can start itiration from d and shift elements arr[i-d] = arr[i]
