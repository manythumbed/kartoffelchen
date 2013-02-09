package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"os"
)

func stave(canvas *pdf.Canvas, origin pdf.Point, width pdf.Unit, seperation pdf.Unit) {
	path := new(pdf.Path)
	increment := pdf.Unit(0)
	for i := 0; i < 5; i++ {
		path.Move(pdf.Point{origin.X, origin.Y + increment})
		path.Line(pdf.Point{origin.X + width, origin.Y + increment})
		canvas.Stroke(path)
		increment = increment + seperation
	}
}

func main() {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.A4Width, pdf.A4Height)
	canvas.Translate(pdf.A4Width/2, pdf.A4Height/2)

	canvas.SetLineWidth(pdf.Unit(125))
	stave(canvas, pdf.Point{50, 50}, pdf.Unit(30), pdf.Unit(2))
	canvas.Close()

	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
