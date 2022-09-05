package main

import "fmt"

func getGroup(value float32) int {
	return int(value/10) * 10 // -25.4 / 10 * 10 = -20
}

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// map with group and temperatures of this group
	m := make(map[int][]float32)

	// loop through temperatures
	for _, t := range temperatures {
		// calculate group of the temperature
		group := getGroup(t)
		// add within the group current temperature
		m[group] = append(m[group], t)
	}

	fmt.Println(m)
}
