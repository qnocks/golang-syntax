package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	// channel to transfer data
	data := make(chan int, len(nums))

	// channel to read multiplied values from
	multipliedNums := make(chan int, len(nums))

	// loop though nums, write to data channel
	for _, num := range nums {
		data <- num
	}

	// don't need the channel to write, so close it
	close(data)

	// read values, convert to multiplied, write to second channel
	for d := range data {
		multipliedNums <- d * 2
	}

	// don't need the channel to write, so close it
	close(multipliedNums)

	// print multiplied values
	for s := range multipliedNums {
		fmt.Println(s)
	}
}
