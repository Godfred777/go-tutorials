package area

type Rectangle struct {
	Length float64
	Breath float64
}

func AreaOfRectangle(length float64, breath float64) float64 {

	rectangle := new(Rectangle)
	rectangle.Length = length
	rectangle.Breath = breath

	area := rectangle.Length * rectangle.Breath
	return area
}
