package methods

import (
	"compmath4/qarrale/src/model"
	"math"
)

type Deg struct {
	A float64
	B float64
}

func (d *Deg) F(x float64) float64 {
	if x > 0 {
		return d.A * math.Pow(x, d.B)
	} else {
		return 0
	}
}

func CompleteDeg(data model.Data) (float64, float64) {

	for i := 0; i < data.Node; i++ {
		data.Nodes_y[i] = math.Log(data.Nodes_y[i])
	}

	for i := 0; i < data.Node; i++ {
		data.Nodes_x[i] = math.Log(data.Nodes_x[i])
	}

	b, a, _ := CompleteLine(data)

	for i := 0; i < data.Node; i++ {
		data.Nodes_y[i] = math.Pow(math.E, data.Nodes_y[i])
	}

	for i := 0; i < data.Node; i++ {
		data.Nodes_x[i] = math.Pow(math.E, data.Nodes_x[i])
	}

	a = math.Pow(math.E, a)
	return a, b

}
