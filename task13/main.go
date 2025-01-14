package main

import (
	"fmt"
)

func main() {
	//Поменять местами два числа без создания временной переменной.
	a, b := 1, 2
	fmt.Println(SwapByAssignment(a, b))
	fmt.Println(SwapByMathOper(a, b))
	fmt.Println(SwapByBitOper(a, b))
	
}

func SwapByAssignment(a, b int) (int, int){
	a, b = b, a

	return a, b
}

func SwapByMathOper(a, b int) (int, int) {
	a += b
	b = a - b
	a -= b

	return a, b
}

func SwapByBitOper(a, b int) (int, int) { // XOR 
	a ^= b 
	b = a ^ b
	a ^= b

	return a, b
}
