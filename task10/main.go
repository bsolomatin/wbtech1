package main

import "fmt"

func main() {
  //Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
	src := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(collectRanges(src))
}

func collectRanges(src []float64) map[int][]float64 {
	res := make(map[int][]float64)
	for _, val := range src {
		key := int(val / 10) * 10
		res[key] = append(res[key], val)
	}

	return res
}
