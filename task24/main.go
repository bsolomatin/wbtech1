package main

import (
	"fmt"
	"math"
)

type Point[T int | float64] struct {
	x, y T
}

func NewPoint[T int | float64](x, y T) *Point[T] {
	return &Point[T]{
		x : x, 
		y : y,
	}
}

func (p *Point[T]) GetX() T {
	return p.x
}

func (p *Point[T]) GetY() T {
	return p.y
}

func CalcDistance[T int | float64](p1, p2 Point[T]) float64 {
	x1 := float64(p1.x)
	y1 := float64(p1.y)
	x2 := float64(p2.x)
	y2 := float64(p2.y)
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2)) // sqrt((x2 - x1) ^ 2 + (y2 - y1) ^ 2)
}

func main() {
	//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
	p1 := Point[int] {
		x: 1, 
		y: 2,
	}
	p2 := Point[int] {
		x: 4, 
		y: 6,
	}
	fmt.Println("Distance between int points:", CalcDistance(p1, p2)) 
	p3 := Point[float64] {
		x: 1.5, 
		y: 2.5,
	}
	p4 := Point[float64] {
		x: 4.5, 
		y: 6.5,
	}
	fmt.Println("Distance between float64 points:", CalcDistance(p3, p4))
}
