package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
	//go run main.go -str="abcdef"
	str := flag.String("str", "World🎉!", "Input string to check unique")
	flag.Parse()
	fmt.Println(isUniqueString(*str))
}

func isUniqueString(str string) bool {
	str = strings.ToLower(str)
	runeMap := make(map[rune]struct{})
	for _, val := range str {
		runeMap[val] = struct{}{}
	}

	return len(runeMap) == len([]rune(str))
}
