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

type terminalInput struct {
	argc uint8
	name string
}

func (t *terminalInput) execute(argc ...string) {

	fmt.Println("Введите количество нод(>=7)")
	fmt.Println("Затем ноды в формате x1 y1\\n")

	data, _ := input_output.ReadData(os.Stdin)
	if data.Node < 7 {
		fmt.Println("Число нод должно быть боле = 7")
		return
	}

	a, b, r := methods.CompleteLine(data)
	first := methods.Line{A: a, B: b, R: r}

	a0, a1, a2 := methods.CompleteQuadr(data)
	second := methods.Quadr{A: a0, B: a1, C: a2}

	a, b = methods.CompleteExp(data)
	third := methods.Exp{A: a, B: b}

	a, b = methods.CompleteDeg(data)
	forth := methods.Deg{A: a, B: b}

	a, b = methods.CompleteLog(data)
	fifth := methods.Log{A: a, B: b}

	sort.Float64s(data.Nodes_x)
	sort.Float64s(data.Nodes_y)

	graphic.Draw(data, graphic.DrawFunction([]methods.Function{&first, &second, &third, &forth, &fifth}, data.Nodes_x[0]-5, data.Nodes_x[data.Node-1]+5, data.Nodes_y[0]-5, data.Nodes_y[data.Node-1]+5))

	doubler := methods.Check([]methods.Function{&first, &second, &third, &forth, &fifth}, data)

	err1 := input_output.WriteData(os.Stdout, "Линейная:", "a:", util.FloatToString(first.A), "b:", util.FloatToString(first.B), "r:", util.FloatToString(first.R), "sigma:", util.FloatToString(doubler[0]))
	err2 := input_output.WriteData(os.Stdout, "Квадратичная:", "a2:", util.FloatToString(second.A), "a1:", util.FloatToString(second.B), "a0:", util.FloatToString(second.C), "sigma:", util.FloatToString(doubler[1]))
	err3 := input_output.WriteData(os.Stdout, "Экспоненциальная:", "a:", util.FloatToString(third.A), "b:", util.FloatToString(third.B), "sigma:", util.FloatToString(doubler[2]))
	err4 := input_output.WriteData(os.Stdout, "Степенная:", "a:", util.FloatToString(forth.A), "b:", util.FloatToString(forth.B), "sigma:", util.FloatToString(doubler[3]))
	err5 := input_output.WriteData(os.Stdout, "Логорифмическая:", "a:", util.FloatToString(fifth.A), "b:", util.FloatToString(fifth.B), "sigma:", util.FloatToString(doubler[4]))

	util.HandleError(&err1)
	util.HandleError(&err2)
	util.HandleError(&err3)
	util.HandleError(&err4)
	util.HandleError(&err5)

}

func (t *terminalInput) getArgs() uint8 {
	return t.argc
}

func (t *terminalInput) getName() string {
	return t.name
}
