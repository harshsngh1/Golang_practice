package main

import (
	"fmt"
	"sync"
)

func CalcSquare(i int, wg *sync.WaitGroup) {
	fmt.Println(i * i)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		CalcSquare(i, &wg)
	}
	wg.Wait()
}
