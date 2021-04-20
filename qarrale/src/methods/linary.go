package methods

import (
	"compmath4/qarrale/src/model"
	"compmath4/qarrale/src/slau"
	"compmath4/qarrale/src/util"
	"math"
)

type Line struct {
	A float64
	B float64
	R float64
}

func (l *Line) F(x float64) float64 {
	return l.A*x + l.B
}

func CompleteLine(data model.Data) (float64, float64, float64) {
	var a, b, r float64
	var nodes_xx = make([]float64, data.Node)
	var nodes_xy = make([]float64, data.Node)

	for i := 0; i < data.Node; i++ {
		nodes_xx[i] = math.Pow(data.Nodes_x[i], 2)
		nodes_xy[i] = data.Nodes_x[i] * data.Nodes_y[i]
	}

	SX := util.Sum(data.Nodes_x, data.Node)
	SY := util.Sum(data.Nodes_y, data.Node)
	SXX := util.Sum(nodes_xx, data.Node)
	SXY := util.Sum(nodes_xy, data.Node)

	x := slau.Solve([][]float64{{SXX, SX}, {SX, float64(data.Node)}}, []float64{SXY, SY})
	a = x[0]
	b = x[1]
	r = pirson(data)

	return a, b, r

}

func pirson(data model.Data) float64 {
	SX := util.Sum(data.Nodes_x, data.Node)
	SY := util.Sum(data.Nodes_y, data.Node)

	xavg := SX / float64(data.Node)
	yavg := SY / float64(data.Node)

	var up, down1, down2 float64
	for i := 0; i < data.Node; i++ {
		up += (data.Nodes_x[i] - xavg) * (data.Nodes_y[i] - yavg)
		down1 += math.Pow(data.Nodes_x[i]-xavg, 2)
		down2 += math.Pow(data.Nodes_y[i]-yavg, 2)

	}

	return up / math.Sqrt(down1*down2)

}
