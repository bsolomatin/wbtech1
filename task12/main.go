package main

import (
	"fmt"
	"strings"
)

func main() {
	//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
	src := "cat, cat, dog, cat, tree"
	fmt.Println(CreateSet(src))
	
}

func CreateSet(src string) map[string]struct{} {
	res := make(map[string]struct{})
	stringArr := strings.Split(src, ", ")
	for _, val := range stringArr {
		if _, exists := res[val]; !exists {
			res[val] = struct{}{}
		}
	}

	return res
}
