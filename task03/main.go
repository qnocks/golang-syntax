package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	// create buffered channel
	squares := make(chan int, len(nums))
	wg := sync.WaitGroup{}

	// assign number of works to complete
	wg.Add(len(nums))

	for _, num := range nums {
		go func(num int) {
			// write data to the channel
			squares <- num * num

			// decrement works number
			wg.Done()
		}(num)
	}

	// wait all works to complete
	wg.Wait()

	// close the channel on write
	close(squares)

	// loop through the channel and calc result
	var result int
	for c := range squares {
		result += c
	}

	fmt.Println(result)
}
