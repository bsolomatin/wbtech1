package main

import (
	"fmt"
	"sync"
)

func main() {
	//Дана последовательность чисел: 2,4,6,8,10.Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
	var wg sync.WaitGroup
	arr := []int{2, 4, 6, 8, 10}
	sum := 0
	ch := make(chan int, len(arr))
	for _, val := range arr {
		wg.Add(1)
		go CalcSquare(&wg, ch, val)
	}

	wg.Wait()
	close(ch)

	for square := range ch {
		sum += square
	}
	
	fmt.Println(sum)
	
}

func CalcSquare(wg *sync.WaitGroup, output chan <- int, val int) {
	defer wg.Done()
	output <- val * val
}
