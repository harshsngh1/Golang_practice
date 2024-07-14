package interview_prep

import (
	"fmt"
	"sync"
)

func fibonacci(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	wg.Add(1)
	go fibonacci(10, ch, &wg)
	wg.Wait()
	for num := range ch {
		fmt.Println(num)
	}
}
