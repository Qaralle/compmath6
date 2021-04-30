package polynoms

import "compmath4/qarrale/src/model"

type Lagranj struct {
	X    float64
	Data model.Data
}

func (l *Lagranj) F(x float64) float64 {
	var result, composition float64 = 0, 1

	for i := 0; i < l.Data.Node; i++ {
		composition = 1

		for j := 0; j < l.Data.Node; j++ {
			if i != j {
				composition *= (x - l.Data.Nodes_x[j]) / (l.Data.Nodes_x[i] - l.Data.Nodes_x[j])
			}
		}
		result += l.Data.Nodes_y[i] * composition
	}

	return result
}
