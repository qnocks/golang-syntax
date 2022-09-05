package main

import (
	"fmt"
	"strconv"
	"time"
)

// https://www.sobyte.net/post/2021-06/several-ways-to-stop-goroutine-in-golang/

func main() {
	c := make(chan string, 6)
	// channel to be used as a semaphore to handle the closing of the goroutine
	done := make(chan struct{})

	go func() {
		i := 0
		for {
			// listen done channel
			select {
			case c <- strconv.Itoa(i):
			case <-done:
				// close channel to stop reading channel's data
				close(c)
				return
			}

			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		// signalize to stop the goroutine by write to done channel after some time
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	for i := range c {
		fmt.Printf("Received: %s\n", i)
	}
}
