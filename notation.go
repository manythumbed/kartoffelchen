package kartoffelchen

import (
	"github.com/manythumbed/kartoffelchen/rational"
)

type Pitch int

func pitch(value int) Pitch	{
	return Pitch(value)
}

type Primitive interface {
	Pitch() (bool, Pitch)
	Duration() rational.Rational
}

type Event struct {
	Primitive
	Position rational.Rational
}

type Rest struct {
	rational.Rational
}

func rest(upper, lower int) Rest	{
	return Rest{rational.New(upper, lower)}
}

func (r Rest) Pitch() (bool, Pitch) {
	return false, pitch(0)
}

func (r Rest) Duration() rational.Rational {
	return  r.Rational
}

type Note struct {
	pitch    Pitch
	rational.Rational
}

func note(value, upper, lower int)	Note {
	return Note{pitch(value), rational.New(upper, lower)}
}

func (n Note) Pitch() (bool, Pitch) {
	return true, n.pitch
}

func (n Note) Duration() rational.Rational {
	return n.Rational
}

func currentPosition(initial, duration rational.Rational) rational.Rational {
	return rational.Add(initial, duration)
}

func Events(notes []Primitive) []Event {
	events := make([]Event, len(notes))
	p := rational.Zero

	for i, n := range notes {
		events[i] = Event{n, p}
		p = currentPosition(p, n.Duration())
	}

	return events
}
