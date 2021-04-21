package command

import (
	"compmath4/qarrale/src/graphic"
	"compmath4/qarrale/src/input_output"
	"compmath4/qarrale/src/methods"
	"compmath4/qarrale/src/util"
	"fmt"
	"os"
	"sort"
)

type fileInput struct {
	argc uint8
	name string
}

func (t *fileInput) execute(argc ...string) {

	file, err := os.Open(argc[0])
	util.HandleError(&err)
	defer file.Close()

	data, _ := input_output.ReadData(file)

	dx := 0.001
	var test bool = false
	for i := 0; i < data.Node; i++ {
		if data.Nodes_x[i] < 0 {
			test = true
		}
		if data.Nodes_y[i] < 0 {
			test = true
		}
		for j := 0; j < data.Node; j++ {

			if data.Nodes_x[i] == data.Nodes_x[j] {
				data.Nodes_x[i] += dx
				dx += 0.0001
				fmt.Println("Существует несколько точек с одинаковым x")
			}
		}
	}

	a, b, r := methods.CompleteLine(data)
	first := methods.Line{a, b, r}

	a0, a1, a2 := methods.CompleteQuadr(data)
	second := methods.Quadr{a0, a1, a2}

	funcl := []methods.Function{&first, &second}

	third := methods.Exp{0, 0}
	forth := methods.Deg{0, 0}
	fifth := methods.Log{0, 0}

	if !test {
		a, b = methods.CompleteExp(data)
		third := methods.Exp{a, b}

		a, b = methods.CompleteDeg(data)
		forth := methods.Deg{a, b}

		a, b = methods.CompleteLog(data)
		fifth := methods.Log{a, b}
		funcl = []methods.Function{&first, &second, &third, &forth, &fifth}
	}

	sort.Float64s(data.Nodes_x)
	sort.Float64s(data.Nodes_y)

	graphic.Draw(data, graphic.DrawFunction(funcl, data.Nodes_x[0]-5, data.Nodes_x[data.Node-1]+5, data.Nodes_y[0]-5, data.Nodes_y[data.Node-1]+5))

	doubler := methods.Check(funcl, data)

	err1 := input_output.WriteData(os.Stdout, "Линейная:", "a:", util.FloatToString(first.A), "b:", util.FloatToString(first.B), "r:", util.FloatToString(first.R), "sigma:", util.FloatToString(doubler[0]))
	err2 := input_output.WriteData(os.Stdout, "Квадратичная:", "a2:", util.FloatToString(second.A), "a1:", util.FloatToString(second.B), "a0:", util.FloatToString(second.C), "sigma:", util.FloatToString(doubler[1]))
	if !test {
		err3 := input_output.WriteData(os.Stdout, "Экспоненциальная:", "a:", util.FloatToString(third.A), "b:", util.FloatToString(third.B), "sigma:", util.FloatToString(doubler[2]))
		err4 := input_output.WriteData(os.Stdout, "Степенная:", "a:", util.FloatToString(forth.A), "b:", util.FloatToString(forth.B), "sigma:", util.FloatToString(doubler[3]))
		err5 := input_output.WriteData(os.Stdout, "Логорифмическая:", "a:", util.FloatToString(fifth.A), "b:", util.FloatToString(fifth.B), "sigma:", util.FloatToString(doubler[4]))
		util.HandleError(&err3)
		util.HandleError(&err4)
		util.HandleError(&err5)
	}

	util.HandleError(&err1)
	util.HandleError(&err2)

}

func (t *fileInput) getArgs() uint8 {
	return t.argc
}

func (t *fileInput) getName() string {
	return t.name
}
