package function

type Function interface {
	F_(x, y float64) float64
	F_clear(x float64) float64
	C_calculate(x, y float64)
	GetName() string
}
