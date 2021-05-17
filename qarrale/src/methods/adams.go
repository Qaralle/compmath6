package methods

import (
	"compmath5/qarrale/src/model"
	"math"
)

type Adams struct {
	Name    string
	Support Method
	Accur   int
}

func (a *Adams) Calculate(data model.Data) []float64 {
	var first, second, third float64

	yValues := make([]float64, len(data.XValues))

	firstFour := a.Support.Calculate(data)[0:4]
	copy(yValues, firstFour)

	for i := 3; i < len(data.XValues)-1; i++ {

		first = data.Function.F_(data.XValues[i], yValues[i]) - data.Function.F_(data.XValues[i-1], yValues[i-1])

		second = data.Function.F_(data.XValues[i], yValues[i]) - 2*data.Function.F_(data.XValues[i-1], yValues[i-1]) +
			data.Function.F_(data.XValues[i-2], yValues[i-2])

		third = data.Function.F_(data.XValues[i], yValues[i]) - 3*data.Function.F_(data.XValues[i-1], yValues[i-1]) +
			3*data.Function.F_(data.XValues[i-2], yValues[i-2]) - data.Function.F_(data.XValues[i-3], yValues[i-3])

		yValues[i+1] = yValues[i] + data.H*data.Function.F_(data.XValues[i], yValues[i]) + (math.Pow(data.H, 2)/2)*
			first + (5*math.Pow(data.H, 3)/12)*second + (3*math.Pow(data.H, 4)/8)*third
	}

	return yValues
}

func (a *Adams) GetAccur() int {
	return a.Accur
}
