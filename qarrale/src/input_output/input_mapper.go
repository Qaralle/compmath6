package input_output

import (
	"bufio"
	"compmath4/qarrale/src/function"
	"compmath4/qarrale/src/util"
	"fmt"
	"strings"
)

var functions map[string]function.Function

func Init() {
	functions = make(map[string]function.Function)
	functions["sinus"] = &function.Sin{Name: "sin x"}
	functions["linear"] = &function.Linear{Name: "1+x-10"}

}

func Prepare(reader *bufio.Reader) function.Function {
	fmt.Println("Выберите нужную функцию")
	for key, _ := range functions {
		fmt.Println("> " + key)
	}
	text, err := reader.ReadString('\n')
	util.HandleError(&err)
	text = strings.TrimSuffix(text, "\n")

	return functions[text]
}
