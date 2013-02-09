package kartoffelchen

type Octave int

type NoteIndex int

type Pitch struct {
	Octave
	NoteIndex
}

func (p Pitch) absolutePitch() int {
	return int(p.Octave*12) + int(p.NoteIndex)
}

func (p Pitch) Transpose(semitones int) Pitch {
	t := p.absolutePitch() + semitones
	if t < 0 && t % 12 != 0	{
		return Pitch{Octave((t / 12) - 1), NoteIndex((12 + (t % 12)) % 12)}
	}
	return Pitch{Octave(t / 12), NoteIndex(t % 12)}
}
