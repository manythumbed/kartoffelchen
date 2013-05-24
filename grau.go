package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"log"
	"github.com/manythumbed/kartoffelchen/engraving"
	"os"
)

func staves(canvas *pdf.Canvas)	{
	r0 := engraving.NewStaffSpec(engraving.RastralZero)
	r1 := engraving.NewStaffSpec(engraving.RastralOne)
	r2 := engraving.NewStaffSpec(engraving.RastralTwo)
	r3 := engraving.NewStaffSpec(engraving.RastralThree)
	r4 := engraving.NewStaffSpec(engraving.RastralFour)
	r5 := engraving.NewStaffSpec(engraving.RastralFive)
	r6 := engraving.NewStaffSpec(engraving.RastralSix)
	r7 := engraving.NewStaffSpec(engraving.RastralSeven)
	r8 := engraving.NewStaffSpec(engraving.RastralEight)

	origin := pdf.Point{21 * pdf.Cm, 1.5 * pdf.Cm}
	engraving.EngraveStaff(origin, 6 * pdf.Cm, r0.Height(), 0.1 * pdf.Pt, canvas)

	origin.Y = origin.Y + r0.Height()  + 1 * pdf.Cm
	engraving.EngraveStaff(origin, 6 * pdf.Cm, r1.Height(), 0.1 * pdf.Pt, canvas)

	origin.Y = origin.Y + r1.Height()  + 1 * pdf.Cm
	engraving.EngraveStaff(origin, 6 * pdf.Cm, r2.Height(), 0.1 * pdf.Pt, canvas)

	origin.Y = origin.Y + r2.Height()  + 1 * pdf.Cm
	engraving.EngraveStaff(origin, 6 * pdf.Cm, r3.Height(), 0.1 * pdf.Pt, canvas)

	origin.Y = origin.Y + r3.Height()  + 1 * pdf.Cm
	engraving.EngraveStaff(origin, 6 * pdf.Cm, r4.Height(), 0.1 * pdf.Pt, canvas)

	origin.Y = origin.Y + r4.Height()  + 1 * pdf.Cm
	engraving.EngraveStaff(origin, 6 * pdf.Cm, r5.Height(), 0.1 * pdf.Pt, canvas)

	origin.Y = origin.Y + r5.Height()  + 1 * pdf.Cm
	engraving.EngraveStaff(origin, 6 * pdf.Cm, r6.Height(), 0.1 * pdf.Pt, canvas)

	origin.Y = origin.Y + r6.Height()  + 1 * pdf.Cm
	engraving.EngraveStaff(origin, 6 * pdf.Cm, r7.Height(), 0.1 * pdf.Pt, canvas)

	origin.Y = origin.Y + r7.Height()  + 1 * pdf.Cm
	engraving.EngraveStaff(origin, 6 * pdf.Cm, r8.Height(), 0.1 * pdf.Pt, canvas)
}

func main() {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.A4Height, pdf.A4Width)

	engraving.Grid(canvas, 1.7, 1.5, 18, 18, 0.2)

	staves(canvas)

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
