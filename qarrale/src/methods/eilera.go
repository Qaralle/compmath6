package methods

import "compmath5/qarrale/src/model"

type Eiler struct {
	Name  string
	Accur int
}

func (e *Eiler) Calculate(data model.Data) []float64 {
	Yvals := make([]float64, len(data.XValues))
	Yvals[0] = data.Y0

	for i := 0; i < len(data.XValues)-1; i++ {
		Yvals[i+1] = Yvals[i] + data.H*data.Function.F_(data.XValues[i], Yvals[i])
	}

	return Yvals
}

func (e *Eiler) GetAccur() int {
	return e.Accur
}
