package area

type Triangle struct {
	base   float64
	height float64
}

const tri_half float64 = 0.5

func AreaOfTriangle(b float64, h float64) float64 {
	triangle := new(Triangle)
	triangle.base = b
	triangle.height = h

	area := tri_half * triangle.base * triangle.height
	return area
}
