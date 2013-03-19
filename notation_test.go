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

	if p, _ := r.Pitch(); p != false {
		t.Errorf("A rest has no pitch")
	}
}

func TestNote(t *testing.T) {
	n := note(2, 4, 4)

	if n.Duration() != rational.New(4, 4) {
		t.Errorf("Note length should be 2, was given %v", n.Duration())
	}

}

func TestEvents(t *testing.T) {
	if len(events(rational.Zero, []Primitive{})) != 0 {
		t.Errorf("An empty slice of primitives should produce an empty slice of events")
	}

	e := events(rational.Zero, []Primitive{rest(4, 4), note(2, 4, 4), rest(4, 4)})
	if len(e) != 3 {
		t.Errorf("Expected 3 events, received %d", len(e))
	}

	if e[0].Position != rational.New(0, 1) {
		t.Errorf("Expected rational.New of 0, received %v", e[0].Position)
	}
	if e[1].Position != rational.New(4, 4) {
		t.Errorf("Expected rational.New of 1, received %v", e[1].Position)
	}
	if e[2].Position != rational.New(8, 4) {
		t.Errorf("Expected rational.New of 2, received %v", e[2].Position)
	}
}

func TestLine(t *testing.T) {
	a := Line{[]Primitive{rest(1, 4), rest(1, 2), rest(1, 4)}}
	b := Line{[]Primitive{rest(1, 4), rest(1, 2), rest(1, 4)}}
	c := Line{[]Primitive{a, b}}

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

func checkRational(expected, received rational.Rational, t *testing.T) {
	if expected != received {
		t.Errorf("Expected %s, received %s", expected, received)
	}
}
