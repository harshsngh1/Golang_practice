package main

func bubble_sort(arr []int) {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		didSwap := 0
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				didSwap = 1
			}
		}
		if didSwap == 0 {
			break
		}
	}

}

// func main() {
// 	arr := []int{64, 25, 12, 22, 11}
// 	fmt.Println("Input array:", arr)

// 	bubble_sort(arr)
// 	fmt.Println("Sorted array:", arr)
// }
