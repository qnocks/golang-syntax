package main

import "fmt"

// helper func to check whether slice contain the element or not
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	var set []string

	// loop through sequence
	for _, s := range sequence {

		// if element isn't  in the result set
		if !contains(set, s) {
			// add to result set
			set = append(set, s)
		}
	}

	fmt.Println(set)
}
