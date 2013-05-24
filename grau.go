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

func main() {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.A4Height, pdf.A4Width)

	engraving.Grid(canvas, 1.7, 1.5, 18, 18, 0.2)

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
