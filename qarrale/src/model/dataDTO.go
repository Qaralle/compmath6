package model

import "compmath5/qarrale/src/function"

type Data struct {
	Name     string
	Node     int
	Nodes_x  []float64
	Nodes_y  []float64
	S        float64
	X        float64
	Function function.Function
}
