package geometry

func mediate(t float64, a, b complex128) complex128 {
	return a + complex(t, 0)*(b-a)
}
