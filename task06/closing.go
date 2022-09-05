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
			// if channel is close, stop goroutine
			data, isOpen := <-c
			if !isOpen {
				fmt.Println("Channel closed")
				return
			}

			fmt.Println(data)
		}
	}()

	c <- "Hello"
	c <- "world!"

	// close channel to stop goroutine
	close(c)
	time.Sleep(time.Second)
}
