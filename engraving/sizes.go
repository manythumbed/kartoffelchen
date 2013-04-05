package engraving

// Rastral staff heights in mm
const (
	RastralZero  = 9.2
	RastralOne   = 7.9
	RastralTwo   = 7.4
	RastralThree = 7.0
	RastralFour  = 6.5
	RastralFive  = 6.0
	RastralSix   = 5.5
	RastralSeven = 4.8
	RastralEight = 3.7
)

func StaffSpace(staffHeight float32) float32 {
	return staffHeight / 4
}
