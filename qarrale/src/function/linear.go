package function

type Linear struct {
	Name string
}

func (s *Linear) F(x float64) float64 {
	return 1 + x - 10
}
