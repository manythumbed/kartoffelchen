package time

import (
	"github.com/manythumbed/kartoffelchen/rational"
)

type Duration rational.Rational

func NewDuration(a, b int) Duration {
	return Duration(rational.New(a, b))
}

func (d Duration) Add(other Duration) Duration {
	return Duration(rational.Add(rational.Rational(d), rational.Rational(other)))
}

func NoDuration() Duration { return Duration(rational.Zero) }

func (a Duration) Greater(b Duration) bool {
	return rational.Greater(rational.Rational(a), rational.Rational(b))
}

type Position rational.Rational

func NewPosition(a, b int) Position {
	return Position(rational.New(a, b))
}

func (p Position) Add(d Duration) Position {
	return Position(rational.Add(rational.Rational(p), rational.Rational(d)))
}

func Zero() Position { return Position(rational.Zero) }

func (a Position) Less(b Position) bool {
	return rational.Less(rational.Rational(a), rational.Rational(b))
}

type TimeSignature struct {
	upper int
	lower rational.Rational
}

func Signature(upper, lower int) TimeSignature {
	return TimeSignature{upper, rational.New(1, lower)}
}

func (t TimeSignature) DurationOfBar() Duration {
	return Duration(rational.Scale(rational.New(t.upper, 1), t.lower))
}
