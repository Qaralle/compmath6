package command

import (
	"compmath5/qarrale/src/function"
	"compmath5/qarrale/src/graphic"
	"compmath5/qarrale/src/input_output"
	"compmath5/qarrale/src/polynoms"
	"fmt"
	"os"
)

type terminalInput struct {
	argc uint8
	name string
}

func (t *terminalInput) execute(argc ...string) {

	fmt.Println("Введите количество нод(>=7)")
	fmt.Println("Введите точку в которой нужно вычеслить значение")
	fmt.Println("Затем ноды в формате x1 y1\\n")

	data, _ := input_output.ReadData(os.Stdin, false)
	if data.Node < 2 {
		fmt.Println("Число нод должно быть боле = 7")
		return
	}

	first := polynoms.Lagranj{data.X, data}
	result := first.F(data.X)
	fmt.Println(result)

	second := polynoms.Newton{data.X, data}
	result2 := second.F(data.X)
	fmt.Println(result2)

	graphic.Draw(data, graphic.DrawFunction([]function.Function{&first, &second}, data.Nodes_x[0]-0.5, data.Nodes_x[data.Node-1]+0.5, data.Nodes_y[0]-0.5, data.Nodes_y[data.Node-1]+0.5))

}

func (t *terminalInput) getArgs() uint8 {
	return t.argc
}

func (t *terminalInput) getName() string {
	return t.name
}
