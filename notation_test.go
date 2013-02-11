package kartoffelchen

import "testing"

func TestPitchTransposition(t *testing.T) {
	for _, d := range transpositions {
		d.check(t)
	}
}

func TestNoteTransposition(t *testing.T)	{
	n1 := Note{Pitch{4, 0}, 1}
	n2 := Note{Pitch{4, 1}, 1}
	if n1.Transpose(1) != n2	{
		t.Errorf("Incorrect transposition. Expected %v for transposition of %d, received %v", n1, 1, n2)
	}
}

func TestRestTranspositions(t *testing.T)	{
	r1 := Rest{1}
	r2 := r1.Transpose(1)
	if r1 != r2 {
		t.Errorf("Incorrect transposition. Expected %v for transposition of %d, received %v", r1, 1, r2)
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

func TestSeqAndComb(t *testing.T)	{
	n1 := Note{Pitch{4, 0}, 1}
	n2 := Note{Pitch{4, 2}, 1}
	n3 := Note{Pitch{4, 4}, 1}
	r1 := Rest{1}
	r2 := Rest{2}

	s1 := seq{}
	s1.AddNote(n1)
	s1.AddRest(r1)
	s1.AddNote(n2)
	s1.AddRest(r2)
	s1.AddNote(n3)

	t.Errorf("%v", s1)

	c1 := comb{}
	c1.AddNote(n1)
	c1.AddNote(n2)
	c1.AddNote(n3)
	t.Errorf("%v", c1)

	s1.AddComb(c1)
	s1.AddNote(n1)
	t.Errorf("%v", s1)

	c2 := comb{}
	c2.AddSeq(s1)
	c2.AddSeq(s1)
	t.Errorf("%v", c2)
}

/*
func TestTransposer(t *testing.T)	{
	n := Note{Pitch{4, 0}, 1}
	t.Errorf("%v", n.Transposed(1))
	t.Errorf("%v", n.Transposed(-12))
}
*/
