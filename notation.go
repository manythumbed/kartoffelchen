package kartoffelchen

import (
	"fmt"
)

type Octave int

type NoteIndex int

type Pitch struct {
	Octave
	NoteIndex
}

type Transposer interface {
	Transpose(semitones int) Transposer
}

func (p Pitch) absolutePitch() int {
	return int(p.Octave*12) + int(p.NoteIndex)
}

func (p Pitch) Transpose(semitones int) Transposer {
	t := p.absolutePitch() + semitones
	if t < 0 && t%12 != 0 {
		return Pitch{Octave((t / 12) - 1), NoteIndex((12 + (t % 12)) % 12)}
	}
	return Pitch{Octave(t / 12), NoteIndex(t % 12)}
}

type Duration rational

type Note struct {
	Pitch
	Duration
}

func (n Note) Transpose(semitones int) Transposer {
	return Note{n.Pitch.Transpose(semitones).(Pitch), n.Duration}
}

type Rest struct {
	Duration
}

func (r Rest) Transpose(semitones int) Transposer {
	return Rest{r.Duration}
}

type seq struct {
	contents []interface{}
}

func (s seq) String() string {
	return fmt.Sprintf("[seq %v]", s.contents)
}

func (s *seq) AddNote(n Note) {
	s.contents = append(s.contents, n)
}

func (s *seq) AddRest(r Rest) {
	s.contents = append(s.contents, r)
}

func (s *seq) AddComb(c comb) {
	s.contents = append(s.contents, c)
}

func (s seq) Transpose(semitones int) Transposer {
	return seq{transposeSlice(s.contents, semitones)}
}

func transposeSlice(s []interface{}, semitones int) []interface{} {
	t := []interface{}{}
	for _, item := range s {
		if original, ok := item.(Transposer); ok {
			t = append(t, original.Transpose(semitones))
		} else {
			t = append(t, item)
		}
	}
	return t
}

type comb struct {
	contents []interface{}
}

func (c comb) String() string {
	return fmt.Sprintf("[comb %v]", c.contents)
}

func (c *comb) AddNote(n Note) {
	c.contents = append(c.contents, n)
}

func (c *comb) AddRest(r Rest) {
	c.contents = append(c.contents, r)
}

func (c *comb) AddSeq(s seq) {
	c.contents = append(c.contents, s)
}

func (c comb) Transpose(semitones int) Transposer {
	return comb{transposeSlice(c.contents, semitones)}
}

type LogicalPosition rational

var origin = LogicalPosition{0, 1}

type MusicalEvent	struct	{
	LogicalPosition
	event interface{}
}

func (s seq) Events(initial LogicalPosition) []MusicalEvent	{
	events := []MusicalEvent{}
	for _, e := range s.contents	{
		events = append(events, MusicalEvent{origin, e})
	}

	return events
}
