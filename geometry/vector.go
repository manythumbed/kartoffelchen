package geometry

import (
	"math"
)

func mediate(t float64, a, b complex128) complex128 {
	return a + complex(t, 0)*(b-a)
}

func normals(a, b complex128) (complex128, complex128) {
	dx := real(b) - real(a)
	dy := imag(b) - imag(a)

	n1 := complex(-1*dy, dx)
	n2 := complex(dy, -1*dx)

	return n1, n2
}

func magnitude(a, b complex128) float64 {
	dx := real(b) - real(a)
	dy := imag(b) - imag(a)

	return math.Sqrt((dx * dx) + (dy * dy))
}
