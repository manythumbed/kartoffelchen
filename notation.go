package kartoffelchen

import (
	"fmt"
	"github.com/manythumbed/kartoffelchen/pitch"
	"github.com/manythumbed/kartoffelchen/time"
)

type Attributes []string

var Untagged = Attributes{}

// Element is the interface that provides the basic methods used by musical elements.
//
// Pitch returns true if the element is pitched with the associated pitch. If the element is
// unpitched it will return false.
//
// Duration is the duration of the element. An element with no duration should 
// return time.NoDuration().
//
// Events are the musical events that make up the element.
type Element interface {
	Pitch() pitch.Pitch
	Duration() time.Duration
	Events(time.Position) []Event
	Tags() Attributes
}

// Event represents a musical element with an associated position in time.
type Event struct {
	Element
	Position time.Position
}

func (e Event) String() string {
	return fmt.Sprintf("%s-%s", e.Element, e.Position)
}

type Rest struct {
	duration time.Duration
	tags     Attributes
}

func (r Rest) Pitch() pitch.Pitch {
	return pitch.Unpitched
}

func (r Rest) Duration() time.Duration	{
	return r.duration
}

func (r Rest) Tags() Attributes {
	return r.tags
}

func (r Rest) Events(start time.Position) []Event {
	return []Event{Event{r, start}}
}

func (r Rest) String() string {
	return fmt.Sprintf("r(%s)", r.duration)
}

type Note struct {
	pitch    pitch.Pitch
	duration time.Duration
	tags     Attributes
}

func (n Note) Pitch() pitch.Pitch {
	return n.pitch
}

func (n Note) Duration() time.Duration {
	return n.duration
}

func (n Note) Events(start time.Position) []Event {
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

func (l Line) Duration() time.Duration {
	d := time.NoDuration()
	for _, e := range l.elements {
		d = d.Add(e.Duration())
	}

	return d
}

func (l Line) Events(start time.Position) []Event {
	e := []Event{}
	for _, v := range l.elements {
		e = append(e, v.Events(start)...)
		start = start.Add(v.Duration())
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

func (l Stack) Duration() time.Duration {
	d := time.NoDuration()
	for _, v := range l.elements {
		if v.Duration().Greater(d) {
			d = v.Duration()
		}
	}

	return d
}

func (s Stack) Events(start time.Position) []Event {
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
