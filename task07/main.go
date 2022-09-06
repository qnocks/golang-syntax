package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ConcurrentMap struct {
	Mutex sync.Mutex
	Entry map[int]int
}

func (m *ConcurrentMap) Add(key int, value int) {
	// lock resource
	m.Mutex.Lock()

	// process operation
	m.Entry[key] = value

	// unlock resource
	m.Mutex.Unlock()
}

func main() {
	// init random
	rand.NewSource(time.Now().UnixNano())
	m := ConcurrentMap{
		Mutex: sync.Mutex{},
		Entry: make(map[int]int),
	}

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			m.Add(rand.Intn(1000), rand.Intn(1000))
			wg.Done()
		}(&wg)
	}

	wg.Wait()
	fmt.Printf("%#v", m.Entry)
}
