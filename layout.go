package kartoffelchen

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"github.com/manythumbed/kartoffelchen/rational"
	"sort"
)

type Bar struct {
	Number int
	Events []Event
}

type TimeSignature struct {
	upper int
	lower rational.Rational
}

func (t TimeSignature) asRational() rational.Rational {
	return rational.Scale(rational.New(t.upper, 1), t.lower)
}

func signature(upper, lower int) TimeSignature {
	return TimeSignature{upper, rational.New(1, lower)}
}

type events []Event

func (e events) Len() int           { return len(e) }
func (e events) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e events) Less(i, j int) bool { return rational.Less(e[i].Position, e[j].Position) }

func before(a, b rational.Rational) bool {
	return rational.Less(a, b)
}

func Bars(signature TimeSignature, initialPosition rational.Rational, element Primitive) []Bar {
	bars := []Bar{}
	eventList := events(element.Events(initialPosition))
	sort.Sort(eventList)
	limit := rational.Add(initialPosition, signature.asRational())
	bar := Bar{1, []Event{}}
	for _, e := range eventList {
		if !before(e.Position, limit) {
			limit = rational.Add(limit, signature.asRational())
			bars = append(bars, bar)
			bar = Bar{bar.Number + 1, []Event{}}
		}

		fmt.Printf("BAR %v\n", limit)
		fmt.Printf("%v\t%v\n", e.Position, e.Primitive)
	}
	bars = append(bars, bar)

	return bars
}

func output() pdf.Document {
	doc := pdf.New()

	return *doc
}
