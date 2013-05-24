package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"log"
	"github.com/manythumbed/kartoffelchen/engraving"
	"os"
)

const (
	borderWidth pdf.Unit = 1.5 * pdf.Cm
)

func grid(canvas *pdf.Canvas, x, y, w, h, step float64) {
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

func main() {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.A4Height, pdf.A4Width)

	grid(canvas, 1.7, 1.5, 18, 18, 0.5)

	canvas.Close()

	file, err := os.Create("grau.pdf")
	if err != nil {
		log.Fatal(err)
	}

	err = doc.Encode(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file.Close()
}
