package command

import (
	"compmath5/qarrale/src/input_output"
	"errors"
)

var commands map[string]Command

func init() {
	commands = make(map[string]Command)
	//commands["terminal"] = &terminalInput{0, "terminal"}
	commands["file"] = &fileInput{1, "file"}
	commands["terminal"] = &terminalHelperInput{0, "terminal"}
}

func dispatch(splitedCommand []string) error {

	currentCommand, ok := commands[splitedCommand[0]]
	input_output.Init()

	if ok {
		if int(currentCommand.getArgs()) == len(splitedCommand)-1 {
			commands[currentCommand.getName()].execute(splitedCommand[1:]...)
		} else {
			return errors.New("invalid structure")
		}
	} else {
		return errors.New("invalid command")
	}

	return nil
}
