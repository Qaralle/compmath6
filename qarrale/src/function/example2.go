package function

import "math"

type Multi struct {
	Name string
	C    float64
}

func (m *Multi) F_(x, y float64) float64 {
	return x * y
}

func (m *Multi) F_clear(x float64) float64 {
	return m.C * math.Pow(math.E, (x*x)/2)
}

func (m *Multi) C_calculate(x, y float64) {
	m.C = y / math.Pow(math.E, (x*x)/2)
}

func (s *Multi) GetName() string {
	return s.Name
}
