package input_output

import (
	"compmath5/qarrale/src/model"
	"compmath5/qarrale/src/util"
)

func Calculate(data *model.Data) {
	a := (data.B - data.A) / data.H
	data.XValues = make([]float64, int(a)+1)

	i := 0

	for x := data.A; x <= data.B; x = util.Round(x+data.H, 5) {
		data.XValues[i] = x
		i++
	}
}
