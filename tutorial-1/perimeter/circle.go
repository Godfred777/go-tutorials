package perimeter

type Arc struct {
	radius float64
	angle  float64
}

const PI float64 = 3.1415

func ArcLength(r float64, a float64) float64 {

	arc := new(Arc)
	arc.angle = a
	arc.radius = r

	if arc.angle > 360 {
		return -1
	}

	arclength := 2 * (arc.angle / 360) * PI * arc.radius
	return arclength

}
