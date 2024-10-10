package mathematics

import (
	"math"
)

const RadianPerDegree = math.Pi / 180.0

const DegreePerRadian = 180.0 / math.Pi

func DivideRoundingUp(dividend, divisor int) int {
	return (dividend + divisor - 1) / divisor
}

func Sgn(x float64) float64 {
	if x < 0.0 {
		return -1.0
	} else if x > 0.0 {
		return 1.0
	} else {
		return 0.0
	}
}
