package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

//General struct for operations
type Number struct {
	intValue  int64     // if int64 suitable
	bigValue  *big.Int  // if int64 overflow
	useBigInt bool      // if int64 overflow
}


func NewNumber(value string) *Number {
	intVal, ok := parseInt64(value) //check that it is [math.minInt64, math.maxInt64]
	if ok {
		return &Number{
			intValue:  intVal,
			bigValue:  nil,
			useBigInt: false,
		}
	}

	// if it is not int64
	bigVal := new(big.Int)
	bigVal.SetString(value, 10)
	return &Number{
		intValue:  0,
		bigValue:  bigVal,
		useBigInt: true,
	}
}

func parseInt64(s string) (int64, bool) {
	bigVal := new(big.Int)
	bigVal.SetString(s, 10)

	if bigVal.Cmp(big.NewInt(math.MaxInt64)) > 0 || bigVal.Cmp(big.NewInt(math.MinInt64)) < 0 { //overflow, bigint required
		return 0, false
	}

	return bigVal.Int64(), true
}

func (n *Number) toBigInt() {
	if !n.useBigInt {
		n.bigValue = big.NewInt(n.intValue)
		n.useBigInt = true
	}
}

func (a *Number) Add(b *Number) *Number {
	if !a.useBigInt && !b.useBigInt {
		sum, overflow := addWithOverflowCheck(a.intValue, b.intValue)
		if !overflow { //if a + b less than math.maxInt64
			return NewNumber(fmt.Sprintf("%d", sum))
		}
	}

	a.toBigInt()
	b.toBigInt()
	result := &Number{
		bigValue:  new(big.Int).Add(a.bigValue, b.bigValue),
		useBigInt: true,
	}
	return result
}

func (n *Number) Subtract(other *Number) *Number {
	if !n.useBigInt && !other.useBigInt {
		diff, overflow := subtractWithOverflowCheck(n.intValue, other.intValue)
		if !overflow { //if a - b more than math.minInt64
			return NewNumber(fmt.Sprintf("%d", diff))
		}
	}

	n.toBigInt()
	other.toBigInt()
	result := &Number{
		bigValue:  new(big.Int).Sub(n.bigValue, other.bigValue),
		useBigInt: true,
	}
	return result
}

func (n *Number) Multiply(other *Number) *Number {
	if !n.useBigInt && !other.useBigInt {
		prod, overflow := multiplyWithOverflowCheck(n.intValue, other.intValue)
		if !overflow { //if a * b less than math.maxInt64
			return NewNumber(fmt.Sprintf("%d", prod))
		}
	}

	n.toBigInt()
	other.toBigInt()
	result := &Number{
		bigValue:  new(big.Int).Mul(n.bigValue, other.bigValue),
		useBigInt: true,
	}
	return result
}

func (a *Number) Divide(b *Number) *Number {
	if !a.useBigInt && !b.useBigInt {
		if b.intValue == 0 {
			panic("Division by zero")
		}
		return NewNumber(fmt.Sprintf("%d", a.intValue/b.intValue))
	}

	a.toBigInt()
	b.toBigInt()
	if b.bigValue.Sign() == 0 {
		panic("Division by zero")
	}
	result := &Number{
		bigValue:  new(big.Int).Div(a.bigValue, b.bigValue),
		useBigInt: true,
	}
	return result
}

func addWithOverflowCheck(a, b int64) (int64, bool) {
	if b > 0 && a > math.MaxInt64-b {
		return 0, false // overflow, i.e int64 res is not used
	}
	if b < 0 && a < math.MinInt64-b {
		return 0, false // overflow, i.e int64 res is not used
	}
	return a + b, true
}

func subtractWithOverflowCheck(a, b int64) (int64, bool) {
	if b < 0 && a > math.MaxInt64+b {
		return 0, false // overflow, i.e int64 res is not used
	}
	if b > 0 && a < math.MinInt64+b {
		return 0, false // overflow, i.e int64 res is not used
	}
	return a - b, true
}

func multiplyWithOverflowCheck(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	result := a * b
	if a == result/b {
		return result, true
	}
	return 0, false // overflow, i.e int64 res is not used
}

func (n *Number) String() string { //for console output
	if n.useBigInt {
		return n.bigValue.String()
	}
	return fmt.Sprintf("%d", n.intValue)
}

func main() {
	/*
	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a и b, значение которых > 2^20.
	До 2^63-1 помещается в int64, если больше, то нужно использовать пакет math/big.
	Хоть размер math/big и динамический, он не может быть меньше 32 байт, 
	поэтому имеет смысл проверить, можно ли обойтись int64, и только если происходит переполнение прибегать к bigInt
	Для этого используется структура, хранящая либо int64, либо bigInt, методы которой проверяют переполнение и создают правильный экземпляр
	*/
	a := NewNumber("2")
	b := NewNumber("2")
	sum := a
	fmt.Printf("Sum: %s and %s = %s\n", a.String(), b.String(), sum.String())

	c := NewNumber(strconv.Itoa(math.MaxInt64))
	d := NewNumber("9999999999999999999")
	sub := c.Subtract(d)
	fmt.Printf("Sub: %s and %s = %s\n", c.String(), d.String(), sub.String())

	e := NewNumber(strconv.Itoa(math.MaxInt64))
	f := NewNumber("2")
	div := e.Divide(f)
	fmt.Printf("Div: %s by %s = %s\n", e.String(), f.String(), div.String())

	g := NewNumber(strconv.Itoa(math.MaxInt64) + "1")
	h := NewNumber("5")
	mul := g.Multiply(h)
	fmt.Printf("Mul: %s by %s = %s\n", g.String(), h.String(), mul.String())
}
