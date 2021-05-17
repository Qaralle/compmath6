package command

import (
	"compmath5/qarrale/src/graphic"
	"compmath5/qarrale/src/input_output"
	"compmath5/qarrale/src/methods"
	"compmath5/qarrale/src/util"
	"fmt"
	"math"
	"os"
)

type terminalHelperInput struct {
	argc uint8
	name string
}

func (t terminalHelperInput) execute(argc ...string) {
	var yh, y2h float64

	data, env, _ := input_output.ReadData(os.Stdin, true)

	flag := util.Check(&data)

	if !flag {
		return
	}

	input_output.Calculate(&data)

	yValues := env.Method.Calculate(data)
	data.YValues = yValues

	data.Function.C_calculate(data.XValues[0], yValues[0])

	for i := 0; i < len(data.XValues); i++ {
		if data.XValues[i] == data.XValues[0]+2*data.H {
			yh = yValues[i]
		}

		if math.IsInf(yValues[i], 1) || math.IsInf(yValues[i], -1) {
			fmt.Println("Слишком большое значение функции")
			return
		}

		fmt.Println("x = ", data.XValues[i], "|| y = ", yValues[i], "|| Yт = ", data.Function.F_clear(data.XValues[i]))
	}

	graphic.Draw(data, graphic.DrawFunction(data.Function, data.A-1, data.B+1))

	data.H /= 2
	input_output.Calculate(&data)
	yValues = env.Method.Calculate(data)

	fmt.Println("точность = ", methods.CalculateCap(yh, y2h, env.Method.GetAccur()))

}

func (t *terminalHelperInput) getArgs() uint8 {
	return t.argc
}

func (t *terminalHelperInput) getName() string {
	return t.name
}
