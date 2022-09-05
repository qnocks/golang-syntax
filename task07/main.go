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

	for i := 0; i < 10; i++ {
		m.Add(rand.Intn(1000), rand.Intn(1000))
	}

	fmt.Printf("%#v", m.Entry)
}
