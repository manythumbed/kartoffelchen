package engraving

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
)

// Rastral staff heights in cm
const (
	RastralZero  = 0.92 * pdf.Cm
	RastralOne   = 0.79 * pdf.Cm
	RastralTwo   = 0.74 * pdf.Cm
	RastralThree = 0.70 * pdf.Cm
	RastralFour  = 0.65 * pdf.Cm
	RastralFive  = 0.60 * pdf.Cm
	RastralSix   = 0.55 * pdf.Cm
	RastralSeven = 0.48 * pdf.Cm
	RastralEight = 0.37 * pdf.Cm
)

func StaffSpace(staffHeight float32) float32 {
	return staffHeight / 4
}

type StaffSpec struct	{
	height pdf.Unit
}

func NewStaffSpec(height pdf.Unit) StaffSpec	{
	return StaffSpec{height}
}

func (s StaffSpec) Height() pdf.Unit {
	return s.height
}

func (s StaffSpec) StaffSpace() pdf.Unit	{
	return pdf.Unit(s.height / 4)
}

func (s StaffSpec) IndexOffset(index int) pdf.Unit	{
	offset := pdf.Unit(s.StaffSpace() / 2)
	return  pdf.Unit(index - 1) * offset
}
