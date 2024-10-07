package main

import (
	"fmt"
	"net/http"
	"sync"
)

func getStatusCodes(url string, wg *sync.WaitGroup, statusChannel chan int) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		statusChannel <- 0
		return
	}
	statusChannel <- response.StatusCode
}

func main() {
	urls := []string{"https://www.google.com/", "https://www.abc123.com/"}
	var wg sync.WaitGroup
	totalUrls := len(urls)
	statusChannel := make(chan int, totalUrls)

	wg.Add(totalUrls)
	for _, url := range urls {
		go getStatusCodes(url, &wg, statusChannel)
	}

	// Closing the channel after all goroutines are done
	// this is done first or before reading from range loop so that channel is closed
	/*
		Yeh goroutine wg.Wait() function call karti hai, jo tab tak block hota hai jab tak saari goroutines jo wg.Add() se add ki gayi thi,
		wg.Done() call karke khatam nahi ho jati. Jab saari goroutines complete ho jaati hain, uske baad hi statusChannel ko close kiya jata hai.
	*/
	go func() {
		wg.Wait()
		close(statusChannel)
	}()

	// Using range to read from the channel
	// this is done after closing so that it does not goes into infinite looping
	for status := range statusChannel {
		fmt.Println(status)
	}
}
