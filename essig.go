package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"log"
	"os"
)

const (
	borderWidth pdf.Unit = 1.5 * pdf.Cm
)

func traceCurve(a, b, c, d pdf.Point, canvas *pdf.Canvas) {
	path := new(pdf.Path)
	path.Move(a)
	path.Curve(b, c, d)

	canvas.Stroke(path)
}

func main() {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.A4Width, pdf.A4Height)

	path := new(pdf.Path)
	bottomLeft := pdf.Point{borderWidth, borderWidth}
	topRight := pdf.Point{pdf.A4Width - borderWidth, pdf.A4Height - borderWidth}

	path.Rectangle(pdf.Rectangle{bottomLeft, topRight})
	canvas.Stroke(path)

	left := 5 * pdf.Cm
	right := 15 * pdf.Cm
	top := 15 * pdf.Cm
	bottom := 10 * pdf.Cm

	curve := new(pdf.Path)
	a := pdf.Point{left, bottom}
	b := pdf.Point{left, top}
	c := pdf.Point{right, top}
	d := pdf.Point{right, bottom}
	curve.Move(a)
	curve.Curve(b, c, d)

	beneath := 5 * pdf.Cm
	e := pdf.Point{left, beneath}
	f := pdf.Point{right, beneath}

	curve.Curve(f, e, a)

	left = 6.5 * pdf.Cm
	right = 13.5 * pdf.Cm
	top = 13.5 * pdf.Cm
	bottom = 10 * pdf.Cm
	beneath = 6.5 * pdf.Cm

	a = pdf.Point{left, bottom}
	b = pdf.Point{left, top}
	c = pdf.Point{right, top}
	d = pdf.Point{right, bottom}
	curve.Move(d)
	curve.Curve(c, b, a)

	e = pdf.Point{left, beneath}
	f = pdf.Point{right, beneath}

	curve.Curve(e, f, d)

	canvas.FillStroke(curve)

	canvas.Close()

	file, err := os.Create("curves.pdf")
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
