package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {
	//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
	//go run main.go -number=1234567 -bit=5
	number := flag.Int64("number", 2, "Input int64 number to reverse i bit")
	bit := flag.Int("bit", 5, "Input bit for reverse. Must be in range [0, 63]")
	flag.Parse()
	if *bit < 0 || *bit > 63 {
		fmt.Println("Bit has invalid value, changed to default 5")
		*bit = 5
	}

	res, err := reverseBit(*number, *bit)
	if err != nil { //i.e irrelevant bit for this number
		fmt.Printf("Fail to change %d bit in %d: %s", *bit, *number, err)
		return
	}

	fmt.Printf("Original number %d, bit to change %d, result %d", *number, *bit, res)

}

func reverseBit(number int64, bitPosition int) (int64, error) {
	bitsCnt := 0
	for n := number; n > 0; n >>= 1 {
		bitsCnt++
	}

	if bitPosition >= bitsCnt { //to set only relevenat bit
		return number, errors.New("bit position must be less than the number of bits in the number")
	}

	return number ^ (1 << bitPosition), nil
}
