package function

import "math"

type Sinx struct {
	Name string
	C    float64
}

func (m *Sinx) F_(x, y float64) float64 {
	return math.Sin(x)
}

func (m *Sinx) F_clear(x float64) float64 {
	return m.C - math.Cos(x)
}

func (m *Sinx) C_calculate(x, y float64) {
	m.C = y + math.Cos(x)
}

func (s *Sinx) GetName() string {
	return s.Name
}
