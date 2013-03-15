package kartoffelchen

type Pitch int
type Duration int
type Position int

type Primitive interface	{
	Pitch() (bool, Pitch)
	Length() (bool, Duration)
}

type Event struct {
	Primitive
	Position
}

type Rest struct	{
	duration Duration
}

func (r Rest) Pitch() (bool, Pitch)	{
	return false, Pitch(0)
}

func (r Rest) Length() (bool, Duration)	{
	return true, r.duration
}

type Note struct	{
	pitch Pitch
	duration Duration
}

func (n Note) Pitch() (bool, Pitch)	{
	return true, n.pitch
}

func (n Note) Length() (bool, Duration)	{
	return true, n.duration
}

func position(initial Position, duration Duration) Position	{
	return Position(int(initial) + int(duration))
}

func Events(notes []Primitive) []Event	{
	events := make([]Event, len(notes))
	p := Position(0)

	for i, n := range notes	{
		events[i] = Event{n, p}
		if ok, dur := n.Length(); ok	{
			p = position(p, dur)
		}
	}

	return events
}
