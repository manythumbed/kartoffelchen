package kartoffelchen

import (
	"github.com/manythumbed/kartoffelchen/pitch"
	"github.com/manythumbed/kartoffelchen/rational"
	"testing"
)

func TestBars(t *testing.T) {
	c := pitch.New(4, 0)
	cs := pitch.New(4, 1)
	qn := rational.New(1, 4)
	hn := rational.New(1, 2)
	none := Untagged

	voice := NewLine(none,
		Note{c, qn, none},
		Note{cs, hn, none},
		Rest{hn, none},
		Note{c, qn, none},
	)
	bars := Bars(signature(3, 4), rational.Zero, voice)

	if len(bars) != 2 {
		t.Errorf("Expected %d bars received %d", 2, len(bars))
	}

	bar1 := Bar{1, []Event{
		Event{voice.elements[0], rational.Zero},
		Event{voice.elements[1], rational.New(1, 4)},
	}}

	bar2 := Bar{2, []Event{
		Event{voice.elements[2], rational.Zero},
		Event{voice.elements[3], rational.New(1, 4)},
	}}

	checkBar(bars[0], bar1, t)
	checkBar(bars[1], bar2, t)
}

func checkBar(received, expected Bar, t *testing.T) {
	if expected.Number != received.Number {
		t.Errorf("Expected bar number to be %v and received %v", expected, received)
	}
	if len(expected.Events) != len(received.Events) {
		t.Errorf("Expected %d events and received %d events", len(expected.Events), len(received.Events))
	}
}
