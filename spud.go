package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"github.com/manythumbed/kartoffelchen/engraving"
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

	largeStaff := engraving.NewStaffSpec(engraving.RastralZero)
	engraving.EngraveStaff(staffOrigin, 12*pdf.Cm, largeStaff.Height(), 0.1*pdf.Pt, canvas)

	engraving.EngraveSurrogateNoteHead(staffOrigin, largeStaff.StaffSpace(), canvas)

	nextNote := pdf.Point{}
	nextNote.X = pdf.Unit(staffOrigin.X + largeStaff.StaffSpace())
	nextNote.Y = pdf.Unit(staffOrigin.Y + largeStaff.IndexOffset(0))
	engraving.EngraveSurrogateNoteHead(nextNote, largeStaff.StaffSpace(), canvas)

	nextNote.X = pdf.Unit(staffOrigin.X + (2 * largeStaff.StaffSpace()))
	nextNote.Y = pdf.Unit(staffOrigin.Y + largeStaff.IndexOffset(7))
	engraving.EngraveSurrogateNoteHead(nextNote, largeStaff.StaffSpace(), canvas)

	nextNote.X = pdf.Unit(staffOrigin.X + (3 * largeStaff.StaffSpace()))
	nextNote.Y = pdf.Unit(staffOrigin.Y + largeStaff.IndexOffset(-1))
	engraving.EngraveSurrogateNoteHead(nextNote, largeStaff.StaffSpace(), canvas)

	nextNote.X = pdf.Unit(staffOrigin.X + (4 * largeStaff.StaffSpace()))
	nextNote.Y = pdf.Unit(staffOrigin.Y + largeStaff.IndexOffset(2))
	engraving.EngraveSurrogateNoteHead(nextNote, largeStaff.StaffSpace(), canvas)

	smallStaff := engraving.NewStaffSpec(engraving.RastralEight)
	staffOrigin.Y = staffOrigin.Y + 5*pdf.Cm
	engraving.EngraveStaff(staffOrigin, 12*pdf.Cm, smallStaff.Height(), 0.1*pdf.Pt, canvas)
	canvas.Close()

	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
