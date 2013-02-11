package kartoffelchen

import (
	"math"
)

type rational struct {
	num, denom int
}

func add(a, b rational) rational {
	l := lcm(a.denom, b.denom)
	c, d := reduce((a.num*(l/a.denom))+(b.num*(l/b.denom)), l)
	return rational{c, d}
}

func scale(a, b rational) rational {
	c, d := reduce(a.num*b.num, a.denom*b.denom)
	return rational{c, d}
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
