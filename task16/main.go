package main

import (
	"fmt"
	"sort"
)

func main() {
	//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
	arr := []int {6, 3, 4, 5, 0, 9, 8, 7, 2, 1}
	fmt.Println(quicksort(arr))
	sort.Slice(arr, func(i, j int) bool{
		return arr[i] < arr[j]
	}) //sort.Slice implemented quicksort inside
	fmt.Println(arr)
}

func quicksort(arr []int) []int {
	if len(arr) <= 1 { // basic case (empty or 1 elem - do not need to sort it again)
		return arr
	}
	pivot := arr[0] //first element as pivot. Also possible to use any other elem as pivot. For example last or middle
	var left, right []int
	for _, val := range arr[1:] { //from second elem till end
		if val <= pivot { //create subarr with elem which less or equal to pivot
			left = append(left, val)
		} else { //create subarr with elem which greater than pivot
			right = append(right, val)
		}
	}

	//Optionally: Can be processed by goroutine with sync mechanism (for example mutex) 
	arr = append(quicksort(left), pivot) //recursive for subarray + pivot (because it has right position already)
	arr = append(arr, quicksort(right)...) //recursive for subarray after pivot. 

	return arr
}
