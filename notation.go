package kartoffelchen

import (
	"fmt"
	"github.com/manythumbed/kartoffelchen/pitch"
	"github.com/manythumbed/kartoffelchen/rational"
)

type Attributes []string

var Untagged = Attributes{}

// Element is the interface that provides the basic methods used by musical elements.
//
// Pitch returns true if the element is pitched with the associated pitch. If the element is
// unpitched it will return false.
//
// Duration is the duration of the element. An element with no duration should return rational.Zero.
//
// Events are the musical events that make up the element.
type Element interface {
	Pitch() pitch.Pitch
	Duration() rational.Rational
	Events(rational.Rational) []Event
	Tags() Attributes
}

// Event represents a musical element with an associated position in time.
type Event struct {
	Element
	Position rational.Rational
}

func (e Event) String() string {
	return fmt.Sprintf("%s-%s", e.Element, e.Position)
}

type Rest struct {
	duration rational.Rational
	tags     Attributes
}

func rest(upper, lower int) Rest {
	return Rest{rational.New(upper, lower), Attributes{}}
}

func (r Rest) Pitch() pitch.Pitch {
	return pitch.Unpitched
}

func (r Rest) Duration() rational.Rational {
	return r.duration
}

func (r Rest) Tags() Attributes {
	return r.tags
}

func (r Rest) Events(start rational.Rational) []Event {
	return []Event{Event{r, start}}
}

func (r Rest) String() string {
	return fmt.Sprintf("r(%s)", r.duration)
}

type Note struct {
	pitch    pitch.Pitch
	duration rational.Rational
	tags     Attributes
}

func note(octave, index, upper, lower int) Note {
	return Note{pitch.New(octave, index), rational.New(upper, lower), Attributes{}}
}

func (n Note) Pitch() pitch.Pitch {
	return n.pitch
}

func (n Note) Duration() rational.Rational {
	return n.duration
}

func (n Note) Events(start rational.Rational) []Event {
	return []Event{Event{n, start}}
}

func (n Note) Tags() Attributes {
	return n.tags
}

type Line struct {
	elements []Element
	tags     Attributes
}

func (l Line) Pitch() pitch.Pitch {
	return pitch.Unpitched
}

func (l Line) Duration() rational.Rational {
	d := rational.Zero
	for _, e := range l.elements {
		d = rational.Add(d, e.Duration())
	}

	return d
}

func (l Line) Events(start rational.Rational) []Event {
	e := []Event{}
	for _, v := range l.elements {
		e = append(e, v.Events(start)...)
		start = rational.Add(start, v.Duration())
	}

	return e
}

func (l Line) Tags() Attributes {
	return l.tags
}

func NewLine(m Attributes, elements ...Element) Line {
	return Line{elements, m}
}

type Stack struct {
	elements []Element
	tags     Attributes
}

func (l Stack) Pitch() pitch.Pitch {
	return pitch.Unpitched
}

func (l Stack) Duration() rational.Rational {
	d := rational.Zero
	for _, v := range l.elements {
		if rational.Greater(v.Duration(), d) {
			d = v.Duration()
		}
	}

	return d
}

func (s Stack) Events(start rational.Rational) []Event {
	e := []Event{}
	for _, v := range s.elements {
		e = append(e, v.Events(start)...)
	}

	return e
}

func (s Stack) Tags() Attributes {
	return s.tags
}

func NewStack(m Attributes, elements ...Element) Stack {
	return Stack{elements, m}
}

func position(initial, duration rational.Rational) rational.Rational {
	return rational.Add(initial, duration)
}
