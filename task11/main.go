package main

import "fmt"

func main() {
	//Реализовать пересечение двух неупорядоченных множеств.
	set1 := []int {1, 3, 2, 5, 4}
	set2 := []int {9, 5, 8, 1, 10} 
	fmt.Println(findOverlaps(set1, set2))
}

func findOverlaps[T comparable](set1, set2 []T) []T{
	var smallSet, largeSet []T
	overlaps := make(map[T]struct{})
	if (len(set1) < len(set2)) {
		smallSet = set1
		largeSet = set2
	} else {
		smallSet = set2
		largeSet = set1
	}

	for _, val := range smallSet {
		overlaps[val]=struct{}{}
	}
	res := make([]T, 0, len(smallSet))
	for _, val := range largeSet {
		if _, contains := overlaps[val]; contains {
			res = append(res, val)
		}
	}

	return res
}
