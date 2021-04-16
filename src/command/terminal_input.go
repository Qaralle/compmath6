package command

import "fmt"

type terminalInput struct {
	argc uint8
	name string
}

func (t *terminalInput) execute(argc ...string){
	fmt.Println("hello")
	fmt.Println("_________________________")
}

func (t *terminalInput) getArgs() uint8{
	return t.argc
}

func (t *terminalInput) getName() string{
	return t.name
}