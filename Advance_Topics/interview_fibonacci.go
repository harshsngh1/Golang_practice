package main

import "fmt"

func fib(n int, out chan int) {
	a1 := 0
	a2 := 1
	out <- a1
	out <- a2
	var sum int
	for i := 2; i < n; i++ {
		sum = a1 + a2
		a1 = a2
		a2 = sum
		out <- sum
	}
	close(out)

}

func main() {
	var n int = 5
	out := make(chan int)
	go fib(n, out)
	for value := range out {
		fmt.Println(value)
	}
}
