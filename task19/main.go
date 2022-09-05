package main

import (
	"fmt"
)

func reverse(s string) string {

	//convert string to slice of rune
	runeString := []rune(s)

	// loop through slice from start and end and swap elements
	for i, j := 0, len(runeString)-1; i < j; i, j = i+1, j-1 {
		runeString[i], runeString[j] = runeString[j], runeString[i]
	}

	return string(runeString)
}

func main() {
	fmt.Println(reverse("⌠⬤⎜ⰺ⚸◰⠇⎆♈Ⱪ▲⮼"))
	fmt.Println(reverse("👽👶🎃"))
}
