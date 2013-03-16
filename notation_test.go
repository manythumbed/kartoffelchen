package kartoffelchen

import (
	"testing"
	"github.com/manythumbed/kartoffelchen/rational"
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
	if len(Events([]Primitive{})) != 0 {
		t.Errorf("An empty slice of primitives should produce an empty slice of events")
	}

	e := Events([]Primitive{rest(4, 4), note(2, 4, 4), rest(4, 4)})
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
