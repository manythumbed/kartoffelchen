package kartoffelchen

import "testing"

func TestPitchTransposition(t *testing.T) {
	for _, d := range transpositions {
		d.check(t)
	}
}

type transpositionCheck struct {
	original, transposed Pitch
	amount               int
}

func (c transpositionCheck) check(t *testing.T) {
	if x := c.original.Transpose(c.amount); x != c.transposed {
		t.Errorf("Expected pitch %v to be %v after transposition of %d, recieved %v", c.original, c.transposed, c.amount, x)
	}
}

var transpositions = []transpositionCheck{
	transpositionCheck{Pitch{4, 0}, Pitch{4, 0}, 0},
	transpositionCheck{Pitch{4, 0}, Pitch{4, 1}, 1},
	transpositionCheck{Pitch{4, 0}, Pitch{5, 0}, 12},
	transpositionCheck{Pitch{4, 0}, Pitch{3, 0}, -12},
	transpositionCheck{Pitch{4, 0}, Pitch{3, 11}, -1},
	transpositionCheck{Pitch{0, 0}, Pitch{-1, 11}, -1},
	transpositionCheck{Pitch{0, 0}, Pitch{-1, 0}, -12},
	transpositionCheck{Pitch{0, 0}, Pitch{-2, 11}, -13},
	transpositionCheck{Pitch{-2, 11}, Pitch{-3, 10}, -13},
}

func TestMusic(t * testing.T)	{
	m := Music{}
	m.AddNote(Note{Pitch{4, 0}, 1})
	m.AddRest(Rest{1})

	n := Music{}
	n.AddRest(Rest{1})
	m.Add(n)
	m.Combine(n)
	t.Errorf("%v", m)
}
