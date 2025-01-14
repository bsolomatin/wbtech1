package main

import (
	"flag"
	"fmt"
)


func main() {
	//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
  //go run main.go -input=главрыба
	str := flag.String("input", "® ✉ § © ☯ ☭ ? $ £ ¢", "Input string")
	flag.Parse()
	fmt.Println(reverse(*str))
	fmt.Println(swapReverse(*str))
}

func reverse(str string ) string {
	strRune := []rune(str)
	strRuneLen := len(strRune)
	res := make([]rune, 0, strRuneLen)
	for i := strRuneLen - 1; i >= 0; i-- {
		res = append(res, strRune[i])
	}

	return string(res)
}

func swapReverse(str string) string {
	strRune := []rune(str)
	strRuneLen := len(strRune)

	for i, j := 0, strRuneLen - 1; i < j; i, j = i + 1, j - 1 {
		strRune[i], strRune[j] = strRune[j], strRune[i]
	}

	return string(strRune)
}
