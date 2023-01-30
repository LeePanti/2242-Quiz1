package main

import (
	"fmt"
	"math"
)

func square(side float64) (float64, float64) {
	area := math.Pow(side, 2)
	perimeter := side * 4
	return area, perimeter
}

func main() {
	fmt.Println(square(5))
}
