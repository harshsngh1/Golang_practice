package main

import (
	"fmt"
)

func selection_sort(arr []int)
{
	n:= len(arr)
	for i:=0;i<n-2;i++ {
		mini:=i
		for j:=i;j<n-1;j++ {
			if arr[j]<arr[mini]{
				mini = j
			}
		}
		arr[i], arr[mini] = arr[mini], arr[i]
	}
}

func main() {
	arr:=[]{13,46,24,52,20,19}
	selection_sort(arr)
	fmt.Println("Sorted array : ",arr)
}