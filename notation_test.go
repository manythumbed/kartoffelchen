package kartoffelchen

import (
	"testing"
)

func TestRest(t *testing.T) {
	r := rest(4, 4)

	l, d := r.Length()
	if l == false {
		t.Errorf("A rest has length")
	}

	if d != duration(4, 4) {
		t.Errorf("Rest length should be 1, was given %v", d)
	}

	if p, _ := r.Pitch(); p != false {
		t.Errorf("A rest has no pitch")
	}
}

func TestNote(t *testing.T) {
	n := note(2, 4, 4)

	l, d := n.Length()
	if l == false {
		t.Errorf("A note has length")
	}

	if d != duration(4, 4) {
		t.Errorf("Note length should be 2, was given %v", d)
	}

}

func TestEvents(t *testing.T) {
	if len(Events([]Primitive{})) != 0 {
		t.Errorf("An empty slice of primitives should produce an empty slice of events")
	}

	e := Events([]Primitive{rest(4, 4), note(2, 4, 4), rest(4, 4)})
	if len(e) != 3 {
		t.Errorf("Expected 3 events, received %d", len(e))
	}

	if e[0].Position != position(0, 1) {
		t.Errorf("Expected position of 0, received %v", e[0].Position)
	}
	if e[1].Position != position(4, 4) {
		t.Errorf("Expected position of 1, received %v", e[1].Position)
	}
	if e[2].Position != position(8, 4) {
		t.Errorf("Expected position of 2, received %v", e[2].Position)
	}
}
