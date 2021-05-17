package methods

import "compmath5/qarrale/src/model"

type Method interface {
	Calculate(data model.Data) []float64
	GetAccur() int
}
