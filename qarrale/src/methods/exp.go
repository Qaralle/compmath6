package methods

import (
	"compmath4/qarrale/src/model"
	"math"
)

type Exp struct {
	A float64
	B float64
}

func (e *Exp) F(x float64) float64 {
	return e.A * math.Pow(math.E, e.B*x)
}

func CompleteExp(data model.Data) (float64, float64) {

	for i := 0; i < data.Node; i++ {
		data.Nodes_y[i] = math.Log(data.Nodes_y[i])
	}

	a, b, _ := CompleteLine(data)

	for i := 0; i < data.Node; i++ {
		data.Nodes_y[i] = math.Pow(math.E, data.Nodes_y[i])
	}

	b = math.Pow(math.E, b)
	return b, a
}
