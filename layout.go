package kartoffelchen

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"github.com/manythumbed/kartoffelchen/time"
	"sort"
)

type Bar struct {
	Number int
	Events []Event
}

type events []Event

func (e events) Len() int           { return len(e) }
func (e events) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e events) Less(i, j int) bool { return e[i].Position.Less(e[j].Position) }

func Bars(signature time.TimeSignature, start time.Position, element Element) []Bar {
	bars := []Bar{}
	eventList := events(element.Events(start))
	sort.Sort(eventList)

	limit := start.Add(signature.DurationOfBar())
	bar := Bar{1, []Event{}}

	for _, e := range eventList {
		if !e.Position.Less(limit) {
			limit = limit.Add(signature.DurationOfBar())
			bars = append(bars, bar)
			bar = Bar{bar.Number + 1, []Event{}}
		}

		bar.Events = append(bar.Events, e)
	}
	bars = append(bars, bar)

	return bars
}

func output() pdf.Document {
	doc := pdf.New()

	return *doc
}
