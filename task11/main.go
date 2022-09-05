package main

import "fmt"

// helper func to check whether slice contain the element or not
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func intersection(firstSet, secondSet []int) []int {
	var result []int

	// loop through minimal set
	if len(firstSet) < len(secondSet) {
		for _, e := range firstSet {
			// if element is in the other set and not in result set
			if contains(secondSet, e) && !contains(result, e) {
				// add element to result set
				result = append(result, e)
			}
		}
	} else {
		// same logic with secondSet as minimal set
		for _, e := range secondSet {
			if contains(firstSet, e) && !contains(result, e) {
				result = append(result, e)
			}
		}
	}

	return result
}

func main() {
	first := []int{1, 2, 3, 4, 5}
	second := []int{3, 7, 1}

	fmt.Println(intersection(first, second))
}
