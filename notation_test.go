package kartoffelchen

import (
	"github.com/manythumbed/kartoffelchen/pitch"
	"github.com/manythumbed/kartoffelchen/time"
	"testing"
)

var (
	c  = pitch.New(4, 0)
	qn = time.NewDuration(1, 4)
	hn = time.NewDuration(1, 2)
	wn = time.NewDuration(4, 4)
)

func TestRest(t *testing.T) {
	r := Rest{wn, Untagged}

	if r.Duration() != wn {
		t.Errorf("Rest length should be 1, was given %v", r.Duration())
	}

	if r.Pitch().Pitched() != false {
		t.Errorf("A rest has no pitch")
	}
}

func TestNote(t *testing.T) {
	n := Note{c, wn, Untagged}

	if n.Duration() != wn {
		t.Errorf("Note length should be 2, was given %v", n.Duration())
	}

}

func TestLine(t *testing.T) {
	r1 := Rest{qn, Untagged}
	r2 := Rest{hn, Untagged}

	a := Line{[]Element{r1, r2, r1}, Attributes{}}
	b := Line{[]Element{r1, r2, r1}, Attributes{}}
	c := Line{[]Element{a, b}, Attributes{}}

	e := c.Events(time.Zero())
	if l := len(e); l != 6 {
		t.Errorf("Expected 6 events, received %v", l)
	}

	checkRational(e[0].Position, time.Zero(), t)
	checkRational(e[1].Position, time.NewPosition(1, 4), t)
	checkRational(e[2].Position, time.NewPosition(3, 4), t)
	checkRational(e[3].Position, time.NewPosition(1, 1), t)
	checkRational(e[4].Position, time.NewPosition(5, 4), t)
	checkRational(e[5].Position, time.NewPosition(7, 4), t)
}

func TestStack(t *testing.T) {
	r1 := Rest{qn, Untagged}
	r2 := Rest{hn, Untagged}

	a := Stack{[]Element{r1, r2, r1}, Attributes{}}
	b := Stack{[]Element{r1, r2, r1}, Attributes{}}
	c := Stack{[]Element{a, b}, Attributes{}}

	e := c.Events(time.Zero())
	if l := len(e); l != 6 {
		t.Errorf("Expected 6 events, received %v", l)
	}

	checkRational(e[0].Position, time.Zero(), t)
	checkRational(e[1].Position, time.Zero(), t)
	checkRational(e[2].Position, time.Zero(), t)
	checkRational(e[3].Position, time.Zero(), t)
	checkRational(e[4].Position, time.Zero(), t)
	checkRational(e[5].Position, time.Zero(), t)
}

func checkRational(received, expected time.Position, t *testing.T) {
	if expected != received {
		t.Errorf("Expected %s, received %s", expected, received)
	}
}
