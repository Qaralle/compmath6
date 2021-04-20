package methods

import (
	"compmath4/qarrale/src/model"
	"math"
)

type Log struct {
	A float64
	B float64
}

func (l *Log) F(x float64) float64 {
	return l.A*math.Log(x) + l.B
}

func CompleteLog(data model.Data) (float64, float64) {

	for i := 0; i < data.Node; i++ {
		data.Nodes_x[i] = math.Log(data.Nodes_x[i])
	}

	a, b, _ := CompleteLine(data)

	for i := 0; i < data.Node; i++ {
		data.Nodes_x[i] = math.Pow(math.E, data.Nodes_x[i])
	}

	return a, b
}
