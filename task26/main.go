package main

import (
	"fmt"
	"strings"
)

func isUniqueChars(s string) bool {
	m := make(map[string]struct{})

	// loop through string converted to lower register
	for _, c := range strings.ToLower(s) {
		// if key found then chars isn't unique
		_, ok := m[string(c)]
		if ok {
			return false
		}

		// add char as key
		m[string(c)] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(isUniqueChars("abcd"))
	fmt.Println(isUniqueChars("abCdefAaf"))
	fmt.Println(isUniqueChars("aabcd"))
}
