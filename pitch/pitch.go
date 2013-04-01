package pitch

import (
	"math"
)

type Octave int

type NoteIndex int

type Pitch struct {
	pitched bool
	octave Octave
	index  NoteIndex
}

func New(octave, index int) Pitch {
	return Pitch{true, Octave(octave), NoteIndex(abs(index))}
}

func (p Pitch) Octave() Octave	{
	return p.octave
}

func (p Pitch) Index() NoteIndex {
	return p.index
}

func (p Pitch) Pitched() bool	{
	return p.pitched
}

func abs(value int) int {
	return int(math.Abs(float64(value))) % 12
}

type Transposer interface {
	Transpose(semitones int) Transposer
}

func (p Pitch) absolutePitch() int {
	return int(p.octave*12) + int(p.index)
}

func (p Pitch) Transpose(semitones int) Transposer {
	if !p.pitched {
		return p
	}

	t := p.absolutePitch() + semitones
	if t < 0 && t%12 != 0 {
		return Pitch{true, Octave((t / 12) - 1), NoteIndex((12 + (t % 12)) % 12)}
	}
	return Pitch{true, Octave(t / 12), NoteIndex(t % 12)}
}

var Unpitched = Pitch{false, 0, -1}
