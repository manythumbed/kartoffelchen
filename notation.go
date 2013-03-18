package kartoffelchen

import (
	"github.com/manythumbed/kartoffelchen/rational"
)

type Pitch int

func pitch(value int) Pitch {
	return Pitch(value)
}

type Primitive interface {
	Pitch() (bool, Pitch)
	Duration() rational.Rational
	Events(rational.Rational) []Event
}

type Event struct {
	Primitive
	Position rational.Rational
}

type Rest struct {
	duration rational.Rational
}

func rest(upper, lower int) Rest {
	return Rest{rational.New(upper, lower)}
}

func (r Rest) Pitch() (bool, Pitch) {
	return false, pitch(0)
}

func (r Rest) Duration() rational.Rational {
	return r.duration
}

func (r Rest) Events(start rational.Rational) []Event {
	return []Event{Event{r, start}}
}

type Note struct {
	pitch    Pitch
	duration rational.Rational
}

func note(value, upper, lower int) Note {
	return Note{pitch(value), rational.New(upper, lower)}
}

func (n Note) Pitch() (bool, Pitch) {
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

func (l Line) Pitch() (bool, Pitch) {
	return false, pitch(0)
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

func (l Stack) Pitch() (bool, Pitch) {
	return false, pitch(0)
}

func (l Stack) Duration() rational.Rational {
	d := rational.Zero
	for _, p := range l.primitives {
		// TO DO need to implement comparison on Rational and / or max
		d = p.Duration()
	}

	return d
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
