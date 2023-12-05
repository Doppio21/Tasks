package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Реализовать конкурентную запись данных в map.

type ConcurrentMap struct {
	mu    sync.RWMutex
	store map[int]int
}

func New() *ConcurrentMap {
	return &ConcurrentMap{
		store: make(map[int]int),
	}
}

func (m *ConcurrentMap) WriteData() {
	m.mu.RLock()
	defer m.mu.RUnlock()

	key := rand.Int()
	value := rand.Int()
	m.store[key] = value
}

func (m *ConcurrentMap) ReadData() {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.store {
		fmt.Printf("key:%d value:%d\n", k, v)
	}
}

func main() {
	wg := sync.WaitGroup{}
	m := New()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.WriteData()
		}()
	}

	wg.Wait()
	m.ReadData()
}
