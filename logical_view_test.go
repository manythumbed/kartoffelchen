package kartoffelchen

import "testing"

func TestSimpleLogicalView(t *testing.T)	{
	notes:= seq{[]interface{}{
		Note{Pitch{4, 0}, Duration{1, 4}},
		Note{Pitch{4, 0}, Duration{1, 4}},
		Note{Pitch{4, 0}, Duration{1, 4}},
		Note{Pitch{4, 0}, Duration{1, 4}},
	}}

	view := []MusicalEvent{
		{LogicalPosition{1, 4}, Note{Pitch{4, 0}, Duration{1, 4}}},
		{LogicalPosition{1, 2}, Note{Pitch{4, 0}, Duration{1, 4}}},
		{LogicalPosition{3, 4}, Note{Pitch{4, 0}, Duration{1, 4}}},
		{LogicalPosition{4, 4}, Note{Pitch{4, 0}, Duration{1, 4}}},
	}

	if len(notes.Events(origin)) != len(view)	{
		t.Errorf("%v %v", view, notes)
	}
}
