package engraving

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
)

func EngraveStaff(origin pdf.Point, width, height, lineWidth pdf.Unit, canvas *pdf.Canvas) {
	path := new(pdf.Path)

	noteHeight := pdf.Unit(height / 4)
	for i := 0; i < 5; i++ {
		path.Move(origin)
		path.Line(pdf.Point{origin.X + width, origin.Y})
		origin.Y = origin.Y + noteHeight
	}

	canvas.Push()
	canvas.SetLineWidth(lineWidth)
	canvas.Stroke(path)
	canvas.Pop()
}
