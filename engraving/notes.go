package engraving

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
)

func EngraveSurrogateNoteHead(origin pdf.Point, size pdf.Unit, canvas *pdf.Canvas) {
	outline := new(pdf.Path)

	topRight := pdf.Point{origin.X + size, origin.Y + size}
	outline.Rectangle(pdf.Rectangle{origin, topRight})

	mid := pdf.Point{origin.X + pdf.Unit(size/2), origin.Y + pdf.Unit(size/2)}
	midPoints := new(pdf.Path)
	midPoints.Move(pdf.Point{mid.X, origin.Y})
	midPoints.Line(pdf.Point{mid.X, origin.Y + size})
	midPoints.Move(pdf.Point{origin.X, mid.Y})
	midPoints.Line(pdf.Point{origin.X + size, mid.Y})

	canvas.Push()
	canvas.SetColor(0.6, 0.6, 0.6)
	canvas.Fill(outline)
	canvas.SetLineWidth(pdf.Unit(0.1))
	canvas.Stroke(midPoints)
	canvas.Pop()
}
