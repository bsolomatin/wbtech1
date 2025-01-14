package main

import (
	"fmt"
	"sort"
)

func main() {
	//Реализовать бинарный поиск встроенными методами языка.
	arr := []int {34, 5, 8, 56, 2, 9, 12} 
	sort.Ints(arr) //mandatory condition for binary search
	fmt.Println(binarySearch(arr, 56)) //manual implementation
	idx := sort.Search(len(arr), func(i int) bool { //sort.Search implements binary search
		return arr[i] >= 56 //to found first relevant element. use == insteadof >= get len(arr) value in case if element is not found. Such variant is easier
	})

	if idx < len(arr) && arr[idx] == 56 {
		fmt.Println(idx)
	} else {
		fmt.Println(-1)
	}
}

func binarySearch(arr []int, target int) int {
	if (len(arr) == 0) { //corner case
		return -1
	}
	
    left, right := 0, len(arr)-1 //initial arr range - all elements
    for left <= right {
        mid := left + (right - left) / 2 //preferable option to save from overflow. Mid is element to collect subarray with less length
        if arr[mid] == target {
            return mid 
        }
        if arr[mid] < target { //check where needed element : on the left or on the right
            left = mid + 1 //resize subarray for continue searching
        } else {
            right = mid - 1
        }
    }
    return -1
}
