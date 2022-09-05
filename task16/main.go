package main

import "fmt"

// wrapper func which calls recursive sorting algorithm
func quickSort(values []int) {
	processSort(values, 0, len(values)-1)
}

func processSort(values []int, left, right int) {

	// stop when right item is greater than left one
	if left < right {
		i := left              // left boundary
		pivot := values[right] // pick a pivot

		// loop through slice, find lower element then pivot and swap
		for j := left; j < right; j++ {
			if values[j] <= pivot {
				values[i], values[j] = values[j], values[i]
				i++
			}
		}

		values[i], values[right] = values[right], values[i]

		// recurse sort for left and right subslices
		processSort(values, left, i-1)
		processSort(values, i+1, right)
	}
}

func main() {
	// TODO: comments
	nums := []int{99, 321, 4, 0, -24, 3}
	quickSort(nums)
	fmt.Println(nums)
}
