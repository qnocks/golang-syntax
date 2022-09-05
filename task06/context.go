package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string)
	// context can be used to identify whether the current channel has been closed
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		i := 0
		for {
			// listen context to close
			select {
			case <-ctx.Done():
				c <- "Context has closed"
				return
			default:
				fmt.Println(strconv.Itoa(i))
			}

			i++
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		// signalize to stop the goroutine by cancel context after some time
		time.Sleep(3 * time.Second)
		cancel()
	}()

	fmt.Println(<-c)
}
