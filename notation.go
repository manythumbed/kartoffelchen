package kartoffelchen

import (
	"fmt"
	"github.com/manythumbed/kartoffelchen/pitch"
	"github.com/manythumbed/kartoffelchen/rational"
)

// Primitive is the interface that provides the basic methods used by musical elements.
//
// Pitch returns true if the element is pitched with the associated pitch. If the element is
// unpitched it will return false.
//
// Duration is the duration of the element. An element with no duration should return rational.Zero.
//
// Events are the musical events that make up the element.
type Primitive interface {
	Pitch() (bool, pitch.Pitch)
	Duration() rational.Rational
	Events(rational.Rational) []Event
}

// Event represents a musical element with an associated position in time.
type Event struct {
	Primitive
	Position rational.Rational
}

func (e Event) String() string {
	return fmt.Sprintf("%s-%s", e.Primitive, e.Position)
}

type Rest struct {
	duration rational.Rational
}

func rest(upper, lower int) Rest {
	return Rest{rational.New(upper, lower)}
}

func (r Rest) Pitch() (bool, pitch.Pitch) {
	return false, pitch.Unpitched
}

func (r Rest) Duration() rational.Rational {
	return r.duration
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
}

func note(octave, index, upper, lower int) Note {
	return Note{pitch.New(octave, index), rational.New(upper, lower)}
}

func (n Note) Pitch() (bool, pitch.Pitch) {
	return true, n.pitch
}

func (n Note) Duration() rational.Rational {
	return n.duration
}

func (n Note) Events(start rational.Rational) []Event {
	return []Event{Event{n, start}}
}

type Line struct {
	primitives []Primitive
}

func (l Line) Pitch() (bool, pitch.Pitch) {
	return false, pitch.Unpitched
}

func (l Line) Duration() rational.Rational {
	d := rational.Zero
	for _, p := range l.primitives {
		d = rational.Add(d, p.Duration())
	}

	return d
}

func (l Line) Events(start rational.Rational) []Event {
	e := []Event{}
	for _, p := range l.primitives {
		e = append(e, p.Events(start)...)
		start = rational.Add(start, p.Duration())
	}

	return e
}

type Stack struct {
	primitives []Primitive
}

func (l Stack) Pitch() (bool, pitch.Pitch) {
	return false, pitch.Unpitched
}

func (l Stack) Duration() rational.Rational {
	d := rational.Zero
	for _, p := range l.primitives {
		if rational.Greater(p.Duration(), d) {
			d = p.Duration()
		}
	}

	return d
}

func (s Stack) Events(start rational.Rational) []Event {
	e := []Event{}
	for _, p := range s.primitives {
		e = append(e, p.Events(start)...)
	}

	return e
}

func position(initial, duration rational.Rational) rational.Rational {
	return rational.Add(initial, duration)
}

func events(initialPosition rational.Rational, notes []Primitive) []Event {
	events := make([]Event, len(notes))
	p := initialPosition

	for i, n := range notes {
		events[i] = Event{n, p}
		p = position(p, n.Duration())
	}

	return events
}
