package kartoffelchen

import (
	"math"
)

type rational struct {
	num, denom int
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func reduce(a, b int) (c, d int) {
	g := gcd(a, b)
	return a / g, b / g
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func lcm(a, b int) int {
	return abs(a*b) / gcd(a, b)
}

func add(a, b rational) rational {
	l := lcm(a.denom, b.denom)
	c, d := reduce((a.num*(a.denom/l))+(b.num*(b.denom/l)), l)
	return rational{c, d}
}
