package area

type Sector struct {
	Radius float64
	Angle  float64
}

const PI float64 = 3.1415

func AreaOfSector(radius float64, angle float64) float64 {

	sector := new(Sector)
	sector.Radius = radius
	sector.Angle = angle

	if sector.Angle > 360 {
		return -1
	}
	area := (sector.Angle / 360) * PI * sector.Radius * sector.Radius
	return area
}
