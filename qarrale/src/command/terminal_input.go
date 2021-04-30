package command

import (
	"compmath4/qarrale/src/graphic"
	"compmath4/qarrale/src/input_output"
	"compmath4/qarrale/src/polynoms"
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

	data, _ := input_output.ReadData(os.Stdin)
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

	graphic.Draw(data, graphic.DrawFunction([]polynoms.FunctionInt{&first, &second}, data.Nodes_x[0]-0.5, data.Nodes_x[data.Node-1]+0.5, data.Nodes_y[0]-0.5, data.Nodes_y[data.Node-1]+0.5))

}

func (t *terminalInput) getArgs() uint8 {
	return t.argc
}

func (t *terminalInput) getName() string {
	return t.name
}
