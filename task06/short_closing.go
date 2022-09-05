package main

import (
	"fmt"
	"time"
)

func main() {
	// channel to transfer data
	c := make(chan string)

	go func() {
		for {
			// syntax sugar for checking if the channel is close
			for v := range c {
				fmt.Println(v)
			}
		}
	}()

	c <- "Hello"
	c <- "world!"

	// close channel to stop the goroutine
	close(c)
	time.Sleep(time.Second)
}
