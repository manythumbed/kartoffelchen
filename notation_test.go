package kartoffelchen

import "testing"

func TestPitchTransposition(t *testing.T) {
	for _, d := range transpositions {
		d.check(t)
	}
}

func TestNoteTransposition(t *testing.T) {
	n1 := Note{Pitch{4, 0}, 1}
	n2 := Note{Pitch{4, 1}, 1}
	if n1.Transpose(1) != n2 {
		t.Errorf("Incorrect transposition. Expected %v for transposition of %d, received %v", n1, 1, n2)
	}
}

func TestRestTranspositions(t *testing.T) {
	r1 := Rest{1}
	r2 := r1.Transpose(1)
	if r1 != r2 {
		t.Errorf("Incorrect transposition. Expected %v for transposition of %d, received %v", r1, 1, r2)
	}
}

type pitchTranspositions struct {
	original, transposed Pitch
	amount               int
}

func (c pitchTranspositions) check(t *testing.T) {
	if x := c.original.Transpose(c.amount); x != c.transposed {
		t.Errorf("Expected pitch %v to be %v after transposition of %d, recieved %v", c.original, c.transposed, c.amount, x)
	}
}

var transpositions = []pitchTranspositions{
	pitchTranspositions{Pitch{4, 0}, Pitch{4, 0}, 0},
	pitchTranspositions{Pitch{4, 0}, Pitch{4, 1}, 1},
	pitchTranspositions{Pitch{4, 0}, Pitch{5, 0}, 12},
	pitchTranspositions{Pitch{4, 0}, Pitch{3, 0}, -12},
	pitchTranspositions{Pitch{4, 0}, Pitch{3, 11}, -1},
	pitchTranspositions{Pitch{0, 0}, Pitch{-1, 11}, -1},
	pitchTranspositions{Pitch{0, 0}, Pitch{-1, 0}, -12},
	pitchTranspositions{Pitch{0, 0}, Pitch{-2, 11}, -13},
	pitchTranspositions{Pitch{-2, 11}, Pitch{-3, 10}, -13},
}

func TestSeqTranspositon(t *testing.T) {
	s1 := seq{[]interface{}{
		Note{Pitch{4, 0}, 1},
		Note{Pitch{4, 2}, 1},
		Note{Pitch{4, 4}, 1},
		Rest{1},
	}}

	s2 := seq{[]interface{}{
		Note{Pitch{4, 1}, 1},
		Note{Pitch{4, 3}, 1},
		Note{Pitch{4, 5}, 1},
		Rest{1},
	}}

	if !s2.equal(s1.Transpose(1).(seq)) {
		t.Errorf("Expected seq %v to be %v after transposition of %d, recieved %v", s1, s2, 1, s1.Transpose(1))
	}
}

func (s seq) equal(s1 seq) bool {
	if len(s.contents) == len(s1.contents) {
		for index := range s.contents {
			switch i := s.contents[index].(type) {
			case Note:
				if j, ok := s1.contents[index].(Note); !ok {
					return false
				} else if i != j {
					return false
				}
			case Rest:
				if j, ok := s1.contents[index].(Rest); !ok {
					return false
				} else if i != j {
					return false
				}
			}
		}
		return true
	}

	return false
}
