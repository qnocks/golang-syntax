package main

import "fmt"

// golang way to swap two vars
func swap(first, second int) (int, int) {
	first, second = second, first
	return first, second
}

// swap using addition and subtraction
func swapUsingAddSub(first, second int) (int, int) {
	second = first + second
	first = second - first
	second = second - first
	return first, second
}

func main() {
	a := 1
	b := 2

	fmt.Printf("Before swapping %d, %d\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After first swapping %d, %d\n", a, b)
	a, b = swapUsingAddSub(a, b)
	fmt.Printf("After second swapping %d, %d\n", a, b)
}
