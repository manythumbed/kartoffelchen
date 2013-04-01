package time

import (
	"github.com/manythumbed/kartoffelchen/rational"
	"testing"
)

func TestAddToPosition(t *testing.T) {
	p := Position(rational.New(1, 2))
	d := Duration(rational.New(1, 4))

	if p.Add(d) != Position(rational.New(3, 4)) {
		t.Errorf("Incorrect value for Add on Position")
	}
}

func TestAddToDuration(t *testing.T) {
	a := Duration(rational.New(1, 2))
	b := Duration(rational.New(1, 4))

	if a.Add(b) != Duration(rational.New(3, 4)) {
		t.Errorf("Incorrect value for Add on Duration")
	}
}
