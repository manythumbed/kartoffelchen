package engraving

import (
	"math"
	"testing"
)

func TestNoteSpace(t *testing.T) {
	n := StaffSpace(RastralZero)
	if !equal(0.230, n, 0.0001) {
		t.Errorf("Incorrect note spacing, expected %3.6f received %3.6f", 2.30, n)
	}
}

func equal(a, b, delta float32) bool {
	return math.Abs(float64(a-b)) < float64(delta)
}
