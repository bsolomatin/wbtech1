package main

import (
	"fmt"
	"sync"
)

func main() {
	//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
	arr := []int{2, 4, 6, 8, 10}
	res := make([]int, len(arr))
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i, val := range arr {
		wg.Add(1)
		go func(index, value int) {
			defer wg.Done()
			mu.Lock() //it is not mandatory because there isn't data race (due to unqiue index), but for fixing warning from -race checker we can do something like this
			res[index] = value * value
			mu.Unlock()
		}(i, val)
	}

	wg.Wait()

	for _, square := range res {
		fmt.Println(square)
	}
}
