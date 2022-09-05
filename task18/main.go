package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentCounter struct {
	Value int
	Mutex sync.Mutex
}

func (c *ConcurrentCounter) Incr() {

	// block access of other goroutines to the shared resource
	c.Mutex.Lock()

	// increment counter
	c.Value++

	// unblock access
	c.Mutex.Unlock()
}

func main() {
	c := ConcurrentCounter{}
	done := make(chan struct{})

	go func() {
		for {
			select {
			case _, ok := <-done:
				if !ok {
					break
				}
			default:
			}

			c.Incr()
		}
	}()

	time.Sleep(2 * time.Second)
	close(done)

	fmt.Printf("Counter value = %d\n", c.Value)
}
