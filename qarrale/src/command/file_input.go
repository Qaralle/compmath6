package command

import (
	"compmath4/qarrale/src/input_output"
	"compmath4/qarrale/src/methods"
	"compmath4/qarrale/src/util"
	"fmt"
	"os"
	"strconv"
)

type fileInput struct {
	argc uint8
	name string
}

func (t *fileInput) execute(argc ...string) {
	fmt.Println("hello, " + argc[0])
	fmt.Println("_________________________")

	file, err := os.Open(argc[0])
	util.HandleError(&err)
	defer file.Close()

	x, _ := input_output.ReadData(file)

	i, _ := strconv.Atoi(x.Name)
	dataMutated := methods.Sum(i, 10)
	x.Name = strconv.Itoa(dataMutated)

	fileo, _ := os.Create(argc[1])
	defer fileo.Close()

	err = input_output.WriteData(fileo, "x")
	util.HandleError(&err)

}

func (t *fileInput) getArgs() uint8 {
	return t.argc
}

func (t *fileInput) getName() string {
	return t.name
}
