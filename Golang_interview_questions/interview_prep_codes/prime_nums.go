package main

import (
	"fmt"
	"math"
	"sync"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func worker(start, end int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if isPrime(i) {
			results <- i
		}
	}
}

func main() {
	N := 20
	numWorkers := 4
	results := make(chan int, N)
	var wg sync.WaitGroup
	rangeSize := (N + numWorkers - 1) / numWorkers // for cases when N is not exactly divided by numWorkers
	// we have just increased the range so that none of the numbers are missed

	for i := 0; i < numWorkers; i++ { // each worker is getting assigned a range to check
		start := i*rangeSize + 1
		end := (i + 1) * rangeSize
		if end > N {
			end = N
		}
		wg.Add(1)
		go worker(start, end, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	primes := []int{}
	for prime := range results {
		primes = append(primes, prime)
	}

	fmt.Println("Prime numbers up to:", N, "are:", primes)
}
