package main

import (
	"flag"
	"fmt"
)

func main() {
	//Удалить i-ый элемент из слайса.
  //go run main.go -i=3
	i := flag.Int("i", 3, "Input index")
	slice := []int {1, 2, 3, 4, 5}
	flag.Parse()
	fmt.Println(removeWithOrderSaving(slice, *i))
	fmt.Println(removeWithoutOrderSaving(slice, *i))
}

func removeWithOrderSaving(slice []int, idx int) []int {
	if idx >= len(slice) {
		return slice //nothing to delete
	}
	return append(slice[:idx], slice[idx+1:]...) //prepare new slice without idx element
}

func removeWithoutOrderSaving(slice []int, idx int) []int {
	if idx >= len(slice) {
		return slice //nothing to delete
	}
	slice[idx] = slice[len(slice) - 1] //set last element on deleted index
	return slice[:len(slice) - 1] //return slice without last element
}
