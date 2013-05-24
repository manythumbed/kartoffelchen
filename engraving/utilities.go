package engraving

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"math"
)

func Grid(canvas *pdf.Canvas, x, y, w, h, step float64) {
	canvas.Push()
	canvas.SetColor(0.75, 0.75, 0.75)
	canvas.Translate(unit(x), unit(y))

	rowCount := int(math.Floor(h / step))

	for row := 0; row < rowCount; row++ {
		bottom := float64(row) * step
		top := bottom + step

		start := 0.0
		if row%2 == 0 {
			start += step
		}
		for left := start; left < w; left += (2 * step) {
			right := left + step
			bottomLeft := point(left, bottom)
			topRight := point(right, top)
			path := new(pdf.Path)
			path.Rectangle(pdf.Rectangle{bottomLeft, topRight})
			canvas.Fill(path)
		}
	}

	canvas.Pop()
}

func point(x, y float64) pdf.Point {
	return pdf.Point{unit(x), unit(y)}
}

func unit(f float64) pdf.Unit {
	return pdf.Unit(f) * pdf.Cm
}
