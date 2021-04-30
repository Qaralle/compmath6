package function

import "math"

type Sin struct {
	Name string
}

func (s *Sin) F(x float64) float64 {
	return math.Sin(x)
}
