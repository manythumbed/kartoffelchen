package kartoffelchen

import (
	"fmt"
	"github.com/manythumbed/kartoffelchen/pitch"
	"github.com/manythumbed/kartoffelchen/rational"
)

type MetaData []string

var Untagged = MetaData{}

// Element is the interface that provides the basic methods used by musical elements.
//
// Pitch returns true if the element is pitched with the associated pitch. If the element is
// unpitched it will return false.
//
// Duration is the duration of the element. An element with no duration should return rational.Zero.
//
// Events are the musical events that make up the element.
type Element interface {
	Pitch() (bool, pitch.Pitch)
	Duration() rational.Rational
	Events(rational.Rational) []Event
	Tags() MetaData
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
	tags     MetaData
}

func rest(upper, lower int) Rest {
	return Rest{rational.New(upper, lower), MetaData{}}
}

func (r Rest) Pitch() (bool, pitch.Pitch) {
	return false, pitch.Unpitched
}

func (r Rest) Duration() rational.Rational {
	return r.duration
}

func (r Rest) Tags() MetaData {
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
	tags     MetaData
}

func note(octave, index, upper, lower int) Note {
	return Note{pitch.New(octave, index), rational.New(upper, lower), MetaData{}}
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

func (n Note) Tags() MetaData {
	return n.tags
}

type Line struct {
	elements []Element
	tags     MetaData
}

func (l Line) Pitch() (bool, pitch.Pitch) {
	return false, pitch.Unpitched
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

func (l Line) Tags() MetaData {
	return l.tags
}

func NewLine(m MetaData, elements ...Element) Line {
	return Line{elements, m}
}

type Stack struct {
	elements []Element
	tags     MetaData
}

func (l Stack) Pitch() (bool, pitch.Pitch) {
	return false, pitch.Unpitched
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

func (s Stack) Tags() MetaData {
	return s.tags
}

func NewStack(m MetaData, elements ...Element) Stack {
	return Stack{elements, m}
}

func position(initial, duration rational.Rational) rational.Rational {
	return rational.Add(initial, duration)
}
