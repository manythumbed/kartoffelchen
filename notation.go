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
	if t < 0 && t%12 != 0 {
		return Pitch{Octave((t / 12) - 1), NoteIndex((12 + (t % 12)) % 12)}
	}
	return Pitch{Octave(t / 12), NoteIndex(t % 12)}
}

type Duration int

type Note struct {
	Pitch
	Duration
}

func (n Note) Transpose(semitones int) Note {
	return Note{n.Pitch.Transpose(semitones), n.Duration}
}

type Rest struct {
	Duration
}

func (r Rest) Transpose(semitones int) Rest {
	return Rest{r.Duration}
}

type MusicalEvent interface {
	Pitched() bool
	GetPitch() *Pitch
	GetDuration() Duration
}

func (r Rest) Pitched() bool	{
	return false
}

func (r Rest) GetPitch() *Pitch	{
	return nil
}

func (r Rest) GetDuration() Duration	{
	return r.Duration
}

/*
func (n Note) Pitched() bool {
	return true
}

func (n Note) GetPitch() Pitch {
	return &n.Pitch
}

func (n Note) GetDuration() Duration {
	return n.Duration
}
*/

/*
	Music is a sequence of musical events, sequential composition
	or parallel composition
*/
type Sequential []MusicalEvent
type Parallel []MusicalEvent

type Music struct	{
	music []interface{}
}

func (m *Music) AddNote(e Note)	{
	m.music = append(m.music, e)
}

func (m *Music) AddRest(r Rest)	{
	m.music = append(m.music, r)
}

func (m *Music) Add(music Music)	{
	m.music = append(m.music, music.music...)
}

func (m *Music) Combine(music Music)	{
	m.music = []interface{}{m.music, music.music}
}
