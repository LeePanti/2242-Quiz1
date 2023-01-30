package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	newTriangle := triangle{3, 4}
	want := 6.00
	got := newTriangle.area()

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestPerimeter(t *testing.T) {
	newTriangle := triangle{3, 4}
	want := 12.00
	got := newTriangle.perimeter()

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
