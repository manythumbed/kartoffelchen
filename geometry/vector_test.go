package geometry

import (
	"math"
	"testing"
)

func TestMediate(t *testing.T) {
	if mediate(0.5, complex(0, 0), complex(1, 1)) != complex(0.5, 0.5) {
		t.Errorf("Mediate produced wrong point")
	}
}

func TestNormals(t *testing.T) {
	n1, n2 := normals(complex(0, 0), complex(1, 1))
	if n1 != complex(-1, 1) || n2 != complex(1, -1) {
		t.Errorf("Expected (-1+1j) and (1-1j) got %v %v", n1, n2)
	}
}

func TestMagnitude(t *testing.T) {
	m := magnitude(complex(0, 0), complex(1, 1))
	if m != math.Sqrt(2) {
		t.Errorf("Expected square root of 2, received %v", m)
	}
}
