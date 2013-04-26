package geometry

import (
	"testing"
)

func TestMediate(t *testing.T) {
	if mediate(0.5, complex(0, 0), complex(1, 1)) != complex(0.5, 0.5) {
		t.Errorf("Mediate produced wrong point")
	}
}
