package kartoffelchen

import (
	"github.com/manythumbed/kartoffelchen/rational"
	"testing"
)

func TestRest(t *testing.T) {
	r := rest(4, 4)

	if r.Duration() != rational.New(4, 4) {
		t.Errorf("Rest length should be 1, was given %v", r.Duration())
	}

	if r.Pitch().Pitched() != false {
		t.Errorf("A rest has no pitch")
	}
}

func TestNote(t *testing.T) {
	n := note(4, 2, 4, 4)

	if n.Duration() != rational.New(4, 4) {
		t.Errorf("Note length should be 2, was given %v", n.Duration())
	}

}

func TestLine(t *testing.T) {
	a := Line{[]Element{rest(1, 4), rest(1, 2), rest(1, 4)}, Attributes{}}
	b := Line{[]Element{rest(1, 4), rest(1, 2), rest(1, 4)}, Attributes{}}
	c := Line{[]Element{a, b}, Attributes{}}

	e := c.Events(rational.Zero)
	if l := len(e); l != 6 {
		t.Errorf("Expected 6 events, received %v", l)
	}

	checkRational(e[0].Position, rational.Zero, t)
	checkRational(e[1].Position, rational.New(1, 4), t)
	checkRational(e[2].Position, rational.New(3, 4), t)
	checkRational(e[3].Position, rational.New(1, 1), t)
	checkRational(e[4].Position, rational.New(5, 4), t)
	checkRational(e[5].Position, rational.New(7, 4), t)
}

func TestStack(t *testing.T) {
	a := Stack{[]Element{rest(1, 4), rest(1, 2), rest(1, 4)}, Attributes{}}
	b := Stack{[]Element{rest(1, 4), rest(1, 2), rest(1, 4)}, Attributes{}}
	c := Stack{[]Element{a, b}, Attributes{}}

	e := c.Events(rational.Zero)
	if l := len(e); l != 6 {
		t.Errorf("Expected 6 events, received %v", l)
	}

	checkRational(e[0].Position, rational.Zero, t)
	checkRational(e[1].Position, rational.Zero, t)
	checkRational(e[2].Position, rational.Zero, t)
	checkRational(e[3].Position, rational.Zero, t)
	checkRational(e[4].Position, rational.Zero, t)
	checkRational(e[5].Position, rational.Zero, t)
}

func checkRational(received, expected rational.Rational, t *testing.T) {
	if expected != received {
		t.Errorf("Expected %s, received %s", expected, received)
	}
}
