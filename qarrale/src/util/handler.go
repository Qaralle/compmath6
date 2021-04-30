package util

import (
	"fmt"
	"strconv"
)

func HandleError(err *error) {
	if *err != nil {
		fmt.Println(*err)
	}
}

func Sum(arr []float64, n int) float64 {
	var sum float64
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	return sum
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

//
//func Check(data model.Data) bool {
//
//}
