package kartoffelchen

import "testing"

func TestPitchTransposition(t *testing.T) {
	for _, d := range transpositions {
		if x := d.original.Transpose(d.amount); x != d.transposed {
			t.Errorf("Expected pitch %v to be %v after transposition of %d, recieved %v", d.original, d.transposed, d.amount, x)
		}
	}
}

type tdata struct {
	original, transposed Pitch
	amount               int
}

var transpositions = []tdata{
	tdata{Pitch{4, 0}, Pitch{4, 0}, 0},
	tdata{Pitch{4, 0}, Pitch{4, 1}, 1},
	tdata{Pitch{4, 0}, Pitch{5, 0}, 12},
	tdata{Pitch{4, 0}, Pitch{3, 0}, -12},
	tdata{Pitch{4, 0}, Pitch{3, 11}, -1},
	tdata{Pitch{0, 0}, Pitch{-1, 11}, -1},
	tdata{Pitch{0, 0}, Pitch{-1, 0}, -12},
	tdata{Pitch{0, 0}, Pitch{-2, 11}, -13},
	tdata{Pitch{-2, 11}, Pitch{-3, 10}, -13},
}
