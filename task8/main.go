package main

import (
	"fmt"
	"sync"
)

type MySyncMap[K comparable, V any] struct {
	mu   sync.Mutex
	data map[K]V
}

func New[K comparable, V any]() *MySyncMap[K, V] {
	return &MySyncMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *MySyncMap[K, V]) Set(key K, value V) {
	defer m.mu.Unlock()
	m.mu.Lock()
	m.data[key] = value
}

func (m *MySyncMap[K, V]) Get(key K) (V, bool) {
	defer m.mu.Unlock()
	m.mu.Lock()
	value, ok := m.data[key]

	return value, ok
}

func main() {
	//Реализовать конкурентную запись данных в map
	res := New[int, int]()
	var wg sync.WaitGroup
	var mu sync.Mutex
	const N = 10

	for i := range N {
		wg.Add(1)
		go func(val int) {
			mu.Lock()
			res.Set(i, i)
			mu.Unlock()
			fmt.Printf("Write %d into the map\n", i)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
}
