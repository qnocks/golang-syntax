package main

import "fmt"

func removeAt(values []int, idx int) []int {
	// check bounds
	if idx < 0 || idx >= len(values) {
		return nil
	}

	// to slice with elements before idx add elements after idx
	return append(values[:idx], values[idx+1:]...)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(removeAt(nums, 2))
}
