package rational

import (
	"fmt"
	"math"
)

type Rational struct {
	num, denom int
}

var Zero = Rational{0, 1}

func New(num, denom int) Rational {
	a, b := reduce(num, denom)
	return Rational{a, b}
}

func Add(a, b Rational) Rational {
	l := lcm(a.denom, b.denom)
	c, d := reduce((a.num*(l/a.denom))+(b.num*(l/b.denom)), l)
	return Rational{c, d}
}

func Scale(a, b Rational) Rational {
	c, d := reduce(a.num*b.num, a.denom*b.denom)
	return Rational{c, d}
}

func (r Rational) String() string {
	return fmt.Sprintf("%d/%d", r.num, r.denom)
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

func Max(rationals []Rational) Rational {
	max := Zero

	for _, r := range rationals {
		if greater(r, max) {
			max = r
		}
	}

	return max
}

func greater(a, b Rational) bool {
	l := lcm(a.denom, b.denom)
	return (a.num * (l / a.denom)) > (b.num * (l / b.denom))
}
