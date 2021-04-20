package methods

import (
	"compmath4/qarrale/src/model"
	"compmath4/qarrale/src/slau"
	"compmath4/qarrale/src/util"
	"math"
)

type Quadr struct {
	A float64
	B float64
	C float64
}

func (q *Quadr) F(x float64) float64 {
	return q.A*x*x + q.B*x + q.C
}

func CompleteQuadr(data model.Data) (float64, float64, float64) {

	var nodes_xx = make([]float64, data.Node)
	var nodes_xy = make([]float64, data.Node)
	var nodes_xxy = make([]float64, data.Node)
	var nodes_xxx = make([]float64, data.Node)
	var nodes_xxxx = make([]float64, data.Node)

	for i := 0; i < data.Node; i++ {
		nodes_xx[i] = math.Pow(data.Nodes_x[i], 2)
		nodes_xxx[i] = math.Pow(data.Nodes_x[i], 3)
		nodes_xxxx[i] = math.Pow(data.Nodes_x[i], 4)
		nodes_xy[i] = data.Nodes_x[i] * data.Nodes_y[i]
		nodes_xxy[i] = math.Pow(data.Nodes_x[i], 2) * data.Nodes_y[i]
	}

	SX := util.Sum(data.Nodes_x, data.Node)
	SY := util.Sum(data.Nodes_y, data.Node)
	SXX := util.Sum(nodes_xx, data.Node)
	SXXX := util.Sum(nodes_xxx, data.Node)
	SXXXX := util.Sum(nodes_xxxx, data.Node)
	SXY := util.Sum(nodes_xy, data.Node)
	SXXY := util.Sum(nodes_xxy, data.Node)

	x := slau.Solve([][]float64{{float64(data.Node), SX, SXX}, {SX, SXX, SXXX}, {SXX, SXXX, SXXXX}}, []float64{SY, SXY, SXXY})

	return x[2], x[1], x[0]

}
