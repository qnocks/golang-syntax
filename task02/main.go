package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	// assign number of works to complete
	wg.Add(len(nums))
	for _, num := range nums {
		go func(num int) {
			fmt.Println(num * num)

			// decrement works number
			wg.Done()
		}(num)
	}

	// wait logic to complete in the main goroutine
	wg.Wait()
}
