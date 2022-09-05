package main

import "fmt"

func binarySearch(values []int, e int) int {
	low := 0
	high := len(values) - 1

	// loop though slice from left, right bounders to mid
	for low <= high {
		// get mid element
		mid := (low + high) / 2
		midVal := values[mid]

		// if element not found change bounders to mid
		if midVal < e {
			low = mid + 1
		} else if midVal > e {
			high = mid - 1
		} else {
			return mid // element found
		}
	}

	return -(low + 1) // element not found
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println(binarySearch(nums, 6))
	fmt.Println(binarySearch(nums, 10))
	fmt.Println(binarySearch(nums, 2))
}
