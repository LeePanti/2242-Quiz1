package main

import (
	"fmt"
	"math"
)

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

func (t triangle) perimeter() float64 {
	hyp := math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2))
	return t.base + t.height + hyp
}

func main() {
	myTriangle := triangle{
		base:   3,
		height: 4,
	}
	fmt.Println(myTriangle.area())
	fmt.Println(myTriangle.perimeter())
}
