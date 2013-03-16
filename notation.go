package kartoffelchen

import (
	"github.com/manythumbed/kartoffelchen/rational"
)

type Pitch int
type Duration rational.Rational
type Position rational.Rational

func pitch(value int) Pitch	{
	return Pitch(value)
}

func duration(upper, lower int) Duration	{
	return Duration(rational.New(upper, lower))
}

func position(upper, lower int) Position {
	return Position(rational.New(upper, lower))
}

type Primitive interface {
	Pitch() (bool, Pitch)
	Length() (bool, Duration)
}

type Event struct {
	Primitive
	Position
}

type Rest struct {
	duration Duration
}

func rest(upper, lower int) Rest	{
	return Rest{duration(upper, lower)}
}

func (r Rest) Pitch() (bool, Pitch) {
	return false, pitch(0)
}

func (r Rest) Length() (bool, Duration) {
	return true, r.duration
}

type Note struct {
	pitch    Pitch
	duration Duration
}

func note(value, upper, lower int)	Note {
	return Note{pitch(value), duration(upper, lower)}
}

func (n Note) Pitch() (bool, Pitch) {
	return true, n.pitch
}

func (n Note) Length() (bool, Duration) {
	return true, n.duration
}

func currentPosition(initial Position, duration Duration) Position {
	return Position(rational.Add(rational.Rational(initial), rational.Rational(duration)))
}

func Events(notes []Primitive) []Event {
	events := make([]Event, len(notes))
	p := Position(rational.Zero)

	for i, n := range notes {
		events[i] = Event{n, p}
		if ok, dur := n.Length(); ok {
			p = currentPosition(p, dur)
		}
	}

	return events
}
