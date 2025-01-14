package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
  //Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
	numberCnt := 10
	var wg sync.WaitGroup
	input := make(chan int)
	output := make(chan int)
	wg.Add(1)
	go CalcSquare(&wg, input, output)
	wg.Add(1)
	go DoPrint(&wg, output)
	for i := range numberCnt {
		input <- i
	}
	close (input)
	wg.Wait()
}

func CalcSquare(wg *sync.WaitGroup, input <- chan int, output chan <- int) {
	defer wg.Done()
	for val := range input {
		time.Sleep(1 * time.Second)
		output <- val * val
	}
	close (output)
}

func DoPrint(wg *sync.WaitGroup, output <- chan int) {
	defer wg.Done()
	for val := range output {
		fmt.Println(val)
	}
}
