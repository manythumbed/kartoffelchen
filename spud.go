package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"os"
)

const (
	borderWidth pdf.Unit = 1.5 * pdf.Cm
)

func main() {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.A4Width, pdf.A4Height)

	path := new(pdf.Path)
	bottomLeft := pdf.Point{borderWidth, borderWidth}
	topRight := pdf.Point{pdf.A4Width - borderWidth, pdf.A4Height - borderWidth}

	path.Rectangle(pdf.Rectangle{bottomLeft, topRight})
	canvas.Stroke(path)

	/*
	path := new(pdf.Path)
	path.Move(pdf.Point{0, 0})
	path.Line(pdf.Point{100, 0})
	canvas.Stroke(path)

	text := new(pdf.Text)
	text.SetFont(pdf.Helvetica, 14)
	text.Text("Hello, World!")
	canvas.DrawText(text)
	*/

	canvas.Close()

	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
