package perimeter

type Quadrilatual struct {
	first_length  float64
	second_length float64
	third_length  float64
	forth_length  float64
}

func Perimeter(fi float64, sc float64, th float64, fo float64) float64 {

	quad := new(Quadrilatual)
	quad.first_length = fi
	quad.second_length = sc
	quad.third_length = th
	quad.forth_length = fo

	per := quad.first_length + quad.second_length + quad.third_length + quad.forth_length
	return per
}
