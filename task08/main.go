package main

import "fmt"

func setBit(num int64, pos, val int) int64 {
	// (1 << pos) - mask with 1 at pos, rest - zeros
	if val == 1 {
		// if setting bit to 1 process OR
		return num | (1 << pos)
	}
	// otherwise process AND
	return num &^ (1 << pos)
}

func main() {
	var num int64 = 4              // 100
	fmt.Println(setBit(num, 0, 1)) // 5 = 101
	fmt.Println(setBit(num, 0, 0)) // 4 = 100
	fmt.Println(setBit(num, 1, 1)) // 6 = 110
	fmt.Println(setBit(num, 1, 0)) // 4 = 100
	fmt.Println(setBit(num, 2, 1)) // 4 = 100
	fmt.Println(setBit(num, 2, 0)) // 0 = 000
	fmt.Println(setBit(num, 3, 1)) // 12 = 1100
}
