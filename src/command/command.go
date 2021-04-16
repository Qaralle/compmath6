package command

type Command interface {
	execute(argc ...string)
	getArgs() uint8
	getName() string
}