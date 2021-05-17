package util

import (
	"compmath5/qarrale/src/model"
	"fmt"
	"math"
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

func Round(x float64, prec int) float64 {
	var rounder float64

	pow := math.Pow(10, float64(prec))
	intermed := x * pow

	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func Check(data *model.Data) bool {
	if data.A > data.B {
		buff := data.A
		data.A = data.B
		data.B = buff
	}

	if data.H < 0 {
		data.H = math.Abs(data.H)
	}

	if data.H > (data.B - data.A) {
		fmt.Println("h слишком большой")
		return false
	}

	if data.H == 0 {
		fmt.Println("h = 0!")
		return false
	}

	return true
}
