package main

import (
	"fmt"
	"time"
)

func sleep(value time.Duration) {
	// waits for the duration to elapse
	<-time.After(value)
}

func main() {
	fmt.Println("Before sleep")
	sleep(3 * time.Second)
	fmt.Println("After sleep")
}
