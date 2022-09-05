package main

import (
	"fmt"
)

func main() {
	// TODO: comments
	nums := []int{2, 4, 6, 8, 10}
	data := make(chan int, len(nums))
	multipliedNums := make(chan int, len(nums))

	for _, num := range nums {
		data <- num
	}
	close(data)

	for d := range data {
		multipliedNums <- d * 2
	}
	close(multipliedNums)

	for s := range multipliedNums {
		fmt.Println(s)
	}
}
