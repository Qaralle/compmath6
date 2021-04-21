package methods

import (
	"compmath4/qarrale/src/model"
	"math"
)

func Check(function []Function, data model.Data) [5]float64 {
	var up float64
	var a = [5]float64{}

	for i, f := range function {
		up = 0
		for i := 0; i < data.Node; i++ {
			up += math.Pow(f.F(data.Nodes_x[i])-data.Nodes_y[i], 2)
		}
		a[i] = math.Sqrt(up / float64(data.Node))
	}
	return a
}
