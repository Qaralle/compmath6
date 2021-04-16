package command

import "fmt"

type fileInput struct {
	argc uint8
	name string
}

func (t *fileInput) execute(argc ...string){
	fmt.Println("hello, "+argc[0])
	fmt.Println("_________________________")
}

func (t *fileInput) getArgs() uint8{
	return t.argc
}

func (t *fileInput) getName() string{
	return t.name
}