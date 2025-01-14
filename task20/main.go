package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».
	//go run main.go -input="snow dog sun — sun dog snow"
	str := flag.String("input", "snow dog sun — sun dog snow", "Input string")
	flag.Parse()
	fmt.Println(reverse(*str))
	fmt.Println(swapReverse(*str))
}

func reverse(str string ) string {
	words := strings.Fields(str)
	wordsCnt := len(words)
	var sb strings.Builder
	for i := wordsCnt - 1; i >= 0; i-- {
		sb.WriteString(words[i])
		if i > 0 {
			sb.WriteByte(' ') //add whitespace after each word except last
		}
	}

	return sb.String()
}

func swapReverse(str string) string {
	words := strings.Fields(str)
	wordsCnt := len(words)

	for i, j := 0, wordsCnt - 1; i < j; i, j = i + 1, j - 1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
