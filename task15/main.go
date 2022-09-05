package main

import (
	"fmt"
	"strings"
)

/*
	Memory Leaking Caused by Substrings

	According to Go specification result string (justString) and base string
	involved in a substring expression (v) share one memory block.
	So, when we assign justString length of 100 we lose 924 bytes
	(the same memory block they share couldn't be collected),
	until the package-level variable justString is modified again elsewhere

	To avoid this memory leak we can clone needed slice of string (create new memory block)
*/

var justString string

func createHugeString(length int) string {
	result := ""
	for length > 0 {
		result += "$"
		length--
	}

	return result
}

func someFunc() {
	v := createHugeString(1 << 10) // 1024 bytes
	justString = v[:100]
}

func someCorrectFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	someCorrectFunc()
	fmt.Println(justString)
}
