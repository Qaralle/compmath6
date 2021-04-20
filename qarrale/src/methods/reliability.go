package methods

import (
	"compmath4/qarrale/src/model"
	"math"
)

func Check(function []Function, data model.Data) [5]float64 {
	var up, down1, down2 float64
	var a = [5]float64{}

	for i, f := range function {
		for i := 0; i < data.Node; i++ {
			up += math.Pow(data.Nodes_y[i]-f.F(data.Nodes_x[i]), 2)
			down1 += math.Pow(f.F(data.Nodes_x[i]), 2)
			down2 += f.F(data.Nodes_x[i])
		}
		a[i] = 1 - (up / (down1 - (float64(1/data.Node) * math.Pow(down2, 2))))
	}
	return a
}
