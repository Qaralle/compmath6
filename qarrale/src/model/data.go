package model

import "compmath5/qarrale/src/function"

type Data struct {
	Name     string
	A        float64
	B        float64
	Y0       float64
	H        float64
	XValues  []float64
	YValues  []float64
	Function function.Function
}
