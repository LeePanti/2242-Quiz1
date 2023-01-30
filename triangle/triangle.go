package main

import (
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
