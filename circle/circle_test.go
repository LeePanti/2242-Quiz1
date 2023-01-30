package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	newCircle := circle{1}
	want := 3.141592653589793
	got := newCircle.area()

	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
