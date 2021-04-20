package slau

import (
	"fmt"
	"math"
	"os"
)

type matrix [][]float64
type vector []float64

func Solve(a matrix, b vector) vector {

	index := make([]int, len(a))
	for i := range index {
		index[i] = i
	}

	for i := 0; i < len(a); i++ {
		r := a[i][index[i]]

		if r == 0 {
			var kk int

			for k := i; k < len(a); k++ {
				if math.Abs(a[i][index[k]]) > r {
					kk = k
				}
			}

			if kk > 0 {
				index[i], index[kk] = index[kk], index[i]
			}

			r = a[i][index[i]]
		}

		if r == 0 {
			if b[i] == 0 {
				fmt.Println("система имеет множество решений")
			} else {
				fmt.Println("система не имеет решений")
			}
			os.Exit(1)
		}

		for j := 0; j < len(a[i]); j++ {
			a[i][index[j]] /= r
		}
		b[i] /= r

		for k := i + 1; k < len(a); k++ {
			r = a[k][index[i]]
			for j := 0; j < len(a[i]); j++ {
				a[k][index[j]] = a[k][index[j]] - a[i][index[j]]*r
			}
			b[k] = b[k] - b[i]*r
		}

	}

	x := make(vector, len(b))

	for i := len(a) - 1; i >= 0; i-- {
		x[i] = b[i]

		for j := i + 1; j < len(a); j++ {
			x[i] = x[i] - (x[j] * a[i][index[j]])
		}
	}

	return x
}
