package engraving

import (
	"math"
	"testing"
)

func TestNoteSpace(t *testing.T) {
	n := NoteSpace(RastralZero)
	if !equal(1.84, n, 0.0001)	{
		t.Errorf("Incorrect note spacing, expected %3.6f received %3.6f", 1.84, n)
	}
}

func equal(a, b, delta float32) bool	{
	return math.Abs(float64(a - b)) < float64(delta)
}
