package main

import (
	"fmt"
	"reflect"
)

func main() {
	data := []interface{}{123, "test", true, make(chan int)}

	for _, d := range data {
		// determine type of var using reflect package
		fmt.Println(reflect.TypeOf(d))
	}
}
