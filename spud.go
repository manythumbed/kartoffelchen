package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"github.com/manythumbed/kartoffelchen/engraving"
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

	staffOrigin := pdf.Point{bottomLeft.X + pdf.Unit(10), bottomLeft.Y + pdf.Unit(500)}
	engraving.EngraveStaff(staffOrigin, 12 * pdf.Cm, 0.97 * pdf.Cm, 0.1 * pdf.Pt, canvas)

	engraving.EngraveSurrogateNoteHead(staffOrigin, 0.1 * pdf.Cm, canvas)

	staffOrigin.Y = staffOrigin.Y + 5 * pdf.Cm
	engraving.EngraveStaff(staffOrigin, 12 * pdf.Cm, 0.37 * pdf.Cm, 0.1 * pdf.Pt, canvas)
	canvas.Close()

	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
