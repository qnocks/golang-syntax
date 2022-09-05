package main

import (
	"fmt"
	"strings"
)

func reverseSentence(sentence string) string {
	// split sentence by spaces
	words := strings.Split(sentence, " ")

	// loop through words from start and end, swap elements
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// join reversed words
	return strings.Join(words, " ")
}

func main() {
	sentence := "snow dog sun"
	fmt.Println(reverseSentence(sentence))
}
