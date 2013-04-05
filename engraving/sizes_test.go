package engraving

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"math"
	"testing"
)

func TestNoteSpace(t *testing.T) {
	spec := NewStaffSpec(RastralZero)
	n := spec.StaffSpace()
	if !equal(float32(0.230*pdf.Cm), float32(n), 0.0001) {
		t.Errorf("Incorrect note spacing, expected %3.6f received %3.6f", 0.230, n)
	}
}

func equal(a, b, delta float32) bool {
	return math.Abs(float64(a-b)) < float64(delta)
}
