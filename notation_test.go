package kartoffelchen

import (
	"testing"
)

func TestRest(t *testing.T)	{
	r := Rest{1}

	l, d := r.Length()
	if l == false {
		t.Errorf("A rest has length")
	}

	if d != 1	{
		t.Errorf("Rest length should be 1, was given %v", d)
	}

	if p,_ := r.Pitch(); p != false	{
		t.Errorf("A rest has no pitch")
	}
}

func TestNote(t *testing.T)	{
	n := Note{1, 2}

	l, d := n.Length()
	if l == false {
		t.Errorf("A note has length")
	}

	if d != 2	{
		t.Errorf("Note length should be 2, was given %v", d)
	}

}

func TestEvents(t *testing.T)	{
	if len(Events([]Primitive{})) != 0	{
		t.Errorf("An empty slice of primitives should produce an empty slice of events")
	}

	e := Events([]Primitive{Rest{1}, Note{1, 2}, Rest{1}})
	if len(e) != 3 {
		t.Errorf("Expected 3 events, received %d", len(e))
	}

	if e[0].Position != 0	{
		t.Errorf("Expected position of 0, received %v", e[0].Position)
	}
	if e[1].Position != 1	{
		t.Errorf("Expected position of 1, received %v", e[1].Position)
	}
	if e[2].Position != 3	{
		t.Errorf("Expected position of 3, received %v", e[2].Position)
	}
}
