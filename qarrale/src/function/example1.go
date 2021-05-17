package function

import "math"

type Sin struct {
	Name string
	C    float64
}

func (s *Sin) F_(x, y float64) float64 {
	return math.Sin(x) + y
}

func (s *Sin) F_clear(x float64) float64 {
	return s.C*math.Exp(x) - math.Sin(x)/2 - math.Cos(x)/2
}

func (s *Sin) C_calculate(x, y float64) {
	s.C = (y + math.Sin(x)/2 + math.Cos(x)/2) / math.Exp(x)
}

func (s *Sin) GetName() string {
	return s.Name
}
