package area

type Trapezium struct {
	Height      float64
	ShortLength float64
	LongLength  float64
}

const trap_half float64 = 0.5

func AreaOfTrapezium(h float64, s float64, l float64) float64 {

	trapezium := new(Trapezium)
	trapezium.Height = h
	trapezium.LongLength = l
	trapezium.ShortLength = s

	area := trap_half * trapezium.Height * (trapezium.LongLength + trapezium.ShortLength)
	return area

}
