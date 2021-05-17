package methods

import "math"

func CalculateCap(yh, y2h float64, p int) float64 {
	return (yh - y2h) / (math.Pow(2, float64(p)) - 1)
}
