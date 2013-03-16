package kartoffelchen

type Pitch int
type Duration rational
type Position rational

func pitch(value int) Pitch	{
	return Pitch(value)
}

func duration(upper, lower int) Duration	{
	a, b := reduce(upper, lower)
	return Duration{a, b}
}

func position(upper, lower int) Position {
	a, b := reduce(upper, lower)
	return Position{a, b}
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
	return Position(add(rational(initial), rational(duration)))
}

func Events(notes []Primitive) []Event {
	events := make([]Event, len(notes))
	p := Position(zero)

	for i, n := range notes {
		events[i] = Event{n, p}
		if ok, dur := n.Length(); ok {
			p = currentPosition(p, dur)
		}
	}

	return events
}
