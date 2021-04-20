package command

import (
	"compmath4/qarrale/src/graphic"
	"compmath4/qarrale/src/input_output"
	"compmath4/qarrale/src/methods"
	"compmath4/qarrale/src/util"
	"os"
	"sort"
)

type terminalInput struct {
	argc uint8
	name string
}

func (t *terminalInput) execute(argc ...string) {

	data, _ := input_output.ReadData(os.Stdin)

	a, b, r := methods.CompleteLine(data)
	first := methods.Line{a, b, r}

	a0, a1, a2 := methods.CompleteQuadr(data)
	second := methods.Quadr{a0, a1, a2}

	a, b = methods.CompleteExp(data)
	third := methods.Exp{a, b}

	a, b = methods.CompleteDeg(data)
	forth := methods.Deg{a, b}

	a, b = methods.CompleteLog(data)
	fifth := methods.Log{a, b}

	sort.Float64s(data.Nodes_x)
	sort.Float64s(data.Nodes_y)

	graphic.Draw(data, graphic.DrawFunction([]methods.Function{&first, &second, &third, &forth, &fifth}, data.Nodes_x[0]-5, data.Nodes_x[data.Node-1]+5, data.Nodes_y[0]-5, data.Nodes_y[data.Node-1]+5))

	err1 := input_output.WriteData(os.Stdout, util.FloatToString(first.A), util.FloatToString(first.B), util.FloatToString(first.R))
	err2 := input_output.WriteData(os.Stdout, util.FloatToString(second.A), util.FloatToString(second.B), util.FloatToString(second.C))
	err3 := input_output.WriteData(os.Stdout, util.FloatToString(third.A), util.FloatToString(third.B))
	err4 := input_output.WriteData(os.Stdout, util.FloatToString(forth.A), util.FloatToString(forth.B))
	err5 := input_output.WriteData(os.Stdout, util.FloatToString(fifth.A), util.FloatToString(fifth.B))

	doubler := methods.Check([]methods.Function{&first, &second, &third, &forth, &fifth}, data)
	for _, rr := range doubler {
		_ = input_output.WriteData(os.Stdout, util.FloatToString(rr))
	}

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
