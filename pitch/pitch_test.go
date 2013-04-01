package pitch

import (
	"testing"
)

func TestNew(t *testing.T) {
	if New(4, -1) != New(4, 1) {
		t.Errorf("Should only allow positive note index")
	}
}

func TestTranspose(t *testing.T) {
	p := New(4, 0)
	if p.Transpose(-1) != New(3, 11) {
		t.Errorf("Should transpose correctly")
	}

	if Unpitched.Transpose(1) != Unpitched {
		t.Errorf("Unpitched should always transpose to Unpitched")
	}
}
