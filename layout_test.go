package kartoffelchen

import (
	"github.com/manythumbed/kartoffelchen/rational"
	"testing"
)

func TestBars(t *testing.T) {
	voice := NewLine(Untagged, note(4, 0, 1, 4), note(4, 1, 1, 2), rest(1, 2), note(4, 0, 1, 4))
	t.Errorf("%v", voice)
	bars := Bars(signature(3, 4), rational.Zero, voice)
	if len(bars) != 2 {
		t.Errorf("Expected %d bars received %d", 2, len(bars))
	}

}
