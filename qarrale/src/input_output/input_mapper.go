package input_output

import (
	"bufio"
	"compmath5/qarrale/src/function"
	methods2 "compmath5/qarrale/src/methods"
	"compmath5/qarrale/src/util"
	"fmt"
	"strings"
)

var functions map[string]function.Function
var methods map[string]methods2.Method

func Init() {
	functions = make(map[string]function.Function)
	functions["sinus 1"] = &function.Sin{Name: "sin x + y"}
	functions["multi"] = &function.Multi{Name: "y' = yx"}
	functions["sinus 2"] = &function.Sinx{Name: "y' = sin x"}

	methods = make(map[string]methods2.Method)
	methods["Eilera"] = &methods2.Eiler{Name: "Eilera", Accur: 1}
	methods["Adamsa"] = &methods2.Adams{Name: "Adamsa", Support: methods["Eilera"], Accur: 4}

}

func PrepareFunction(reader *bufio.Reader) function.Function {
	fmt.Println("Выберите нужную функцию")
	for key, val := range functions {
		fmt.Println("> " + key + " || " + val.GetName())
	}
	text, err := reader.ReadString('\n')
	util.HandleError(&err)
	text = strings.TrimSuffix(text, "\n")

	return functions[text]
}

func PrepareMethod(reader *bufio.Reader) methods2.Method {
	fmt.Println("Выберите нужный метод")
	for key, _ := range methods {
		fmt.Println("> " + key)
	}
	text, err := reader.ReadString('\n')
	util.HandleError(&err)
	text = strings.TrimSuffix(text, "\n")

	return methods[text]
}
