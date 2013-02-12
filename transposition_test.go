package kartoffelchen

import "testing"

func TestPitchTransposition(t *testing.T) {
	for _, d := range transpositions {
		d.check(t)
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

func TestNoteTransposition(t *testing.T) {
	n1 := Note{Pitch{4, 0}, Duration{1, 1}}
	n2 := Note{Pitch{4, 1}, Duration{1, 1}}
	if n1.Transpose(1) != n2 {
		t.Errorf("Incorrect transposition. Expected %v for transposition of %d, received %v", n1, 1, n2)
	}
}

func TestRestTranspositions(t *testing.T) {
	r1 := Rest{Duration{1, 1}}
	r2 := r1.Transpose(1)
	if r1 != r2 {
		t.Errorf("Incorrect transposition. Expected %v for transposition of %d, received %v", r1, 1, r2)
	}
}

func TestSeqTranspositon(t *testing.T) {
	s1 := seq{[]interface{}{
		Note{Pitch{4, 0}, Duration{1, 1}},
		Note{Pitch{4, 2}, Duration{1, 1}},
		Note{Pitch{4, 4}, Duration{1, 1}},
		Rest{Duration{1, 1}},
	}}

	s2 := seq{[]interface{}{
		Note{Pitch{4, 1}, Duration{1, 1}},
		Note{Pitch{4, 3}, Duration{1, 1}},
		Note{Pitch{4, 5}, Duration{1, 1}},
		Rest{Duration{1, 1}},
	}}

	if !s2.equalSeq(s1.Transpose(1).(seq)) {
		t.Errorf("Expected seq %v to be %v after transposition of %d, recieved %v", s1, s2, 1, s1.Transpose(1))
	}
}

func TestCombTransposition(t *testing.T) {
	c1 := comb{[]interface{}{
		Note{Pitch{4, 0}, Duration{1, 1}},
		Note{Pitch{4, 2}, Duration{1, 1}},
		Note{Pitch{4, 4}, Duration{1, 1}},
	}}

	c2 := comb{[]interface{}{
		Note{Pitch{4, 1}, Duration{1, 1}},
		Note{Pitch{4, 3}, Duration{1, 1}},
		Note{Pitch{4, 5}, Duration{1, 1}},
	}}

	if !c2.equalComb(c1.Transpose(1).(comb)) {
		t.Errorf("Expected comb %v to be %v after transposition of %d, recieved %v", c1, c2, 1, c1.Transpose(1))
	}
}

func TestMixedTransposition(t *testing.T) {
	m1 := comb{[]interface{}{
		seq{[]interface{}{
			Note{Pitch{4, 1}, Duration{1, 1}},
			Note{Pitch{4, 3}, Duration{1, 1}},
			Note{Pitch{4, 5}, Duration{1, 1}},
		}},
		seq{[]interface{}{
			Note{Pitch{3, 1}, Duration{1, 1}},
			Note{Pitch{3, 3}, Duration{1, 1}},
			Note{Pitch{3, 5}, Duration{1, 1}},
			comb{[]interface{}{Note{Pitch{2, 11}, Duration{1, 1}}, Note{Pitch{5, 2}, Duration{1, 1}}}},
		}},
		seq{[]interface{}{
			Rest{Duration{1, 1}},
		}},
	}}

	m2 := comb{[]interface{}{
		seq{[]interface{}{
			Note{Pitch{4, 0}, Duration{1, 1}},
			Note{Pitch{4, 2}, Duration{1, 1}},
			Note{Pitch{4, 4}, Duration{1, 1}},
		}},
		seq{[]interface{}{
			Note{Pitch{3, 0}, Duration{1, 1}},
			Note{Pitch{3, 2}, Duration{1, 1}},
			Note{Pitch{3, 4}, Duration{1, 1}},
			comb{[]interface{}{Note{Pitch{2, 10}, Duration{1, 1}}, Note{Pitch{5, 1}, Duration{1, 1}}}},
		}},
		seq{[]interface{}{
			Rest{Duration{1, 1}},
		}},
	}}

	if !m2.equalComb(m1.Transpose(-1).(comb)) {
		t.Errorf("Expected comb %v to be %v after transposition of %d, recieved %v", m1, m2, -1, m1.Transpose(-1))
	}
}

func (s seq) equalSeq(s1 seq) bool {
	return equalSlice(s.contents, s1.contents)
}

func (c comb) equalComb(c1 comb) bool {
	return equalSlice(c.contents, c1.contents)
}

func equalSlice(s1, s2 []interface{}) bool {
	if len(s1) == len(s2) {
		for index := range s1 {
			switch i := s1[index].(type) {
			case Note:
				if j, ok := s2[index].(Note); !ok {
					return false
				} else if i != j {
					return false
				}
			case Rest:
				if j, ok := s2[index].(Rest); !ok {
					return false
				} else if i != j {
					return false
				}
			case seq:
				if j, ok := s2[index].(seq); !ok {
					return false
				} else if !i.equalSeq(j) {
					return false
				}
			case comb:
				if j, ok := s2[index].(comb); !ok {
					return false
				} else if !i.equalComb(j) {
					return false
				}
			default:
				return false
			}
		}
		return true
	}

	return false
}
