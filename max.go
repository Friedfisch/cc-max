package ccmax

import (
	"math"
)

func Max(a float64, b float64) float64 {

	var s uint64 = 1 << 63
	a_b := math.Float64bits(a - b)
	b_a := math.Float64bits(b - a)

	bgta := a_b & s >> 63
	agtb := b_a & s >> 63
	tmp := float64(bgta)*b + float64(agtb)*a

	mul := (a_b + s) & (b_a + s)
	mul1 := mul & s
	tmp2 := float64(mul1>>63) * a

	return tmp + tmp2
}

func MaxCompact(a float64, b float64) float64 {
	return float64(math.Float64bits(a-b)&(1<<63)>>63)*b + float64(math.Float64bits(b-a)&(1<<63)>>63)*a +
		float64((math.Float64bits(a-b)+(1<<63))&(math.Float64bits(b-a)+(1<<63))&(1<<63)>>63)*a
}

func MaxCmp(a float64, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}
