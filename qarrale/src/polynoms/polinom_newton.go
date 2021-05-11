package polynoms

import "compmath5/qarrale/src/model"

type Newton struct {
	X    float64
	Data model.Data
}

func (n *Newton) F(x float64) float64 {
	var result, composition float64 = n.Data.Nodes_y[0], 1
	for i := 1; i < n.Data.Node; i++ {
		composition = 1
		for j := 0; j < i; j++ {
			composition *= x - n.Data.Nodes_x[j]
		}
		result += composition * multif(n.Data, 0, i)
	}
	return result
}

func multif(data model.Data, i, k int) float64 {
	var result float64 = 0
	if k-i == 1 {
		result = (data.Nodes_y[i+1] - data.Nodes_y[k-1]) / (data.Nodes_x[k] - data.Nodes_x[i])
		return result
	}
	result = (multif(data, i+1, k) - multif(data, i, k-1)) / (data.Nodes_x[k] - data.Nodes_x[i])
	return result

}
