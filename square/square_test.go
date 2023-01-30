package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	areaGot, perimeterGot := square(5)
	areaWant := 25.00
	perimeterWant := 20.00

	if areaWant != areaGot || perimeterWant != perimeterGot {
		t.Errorf("Area expected is %v, got %v. Perimeter expected is %v, got %v", areaWant, areaGot, perimeterWant, perimeterGot)
	}
}
