package kartoffelchen

import (
	"math"
)

type Octave int

type NoteIndex int

type Pitch struct {
	octave Octave
	index  NoteIndex
}

func NewPitch(octave, index int) Pitch {
	return Pitch{Octave(octave), NoteIndex(abs(index))}
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
	t := p.absolutePitch() + semitones
	if t < 0 && t%12 != 0 {
		return Pitch{Octave((t / 12) - 1), NoteIndex((12 + (t % 12)) % 12)}
	}
	return Pitch{Octave(t / 12), NoteIndex(t % 12)}
}

var Unpitched = Pitch{0, -1}
