package command

import (
	"bufio"
	"compmath4/qarrale/src/util"
	"fmt"
	"os"
	"strings"
)

func Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>>")
		text, err := reader.ReadString('\n')
		util.HandleError(&err)
		if text != "" {
			text = strings.TrimSuffix(text, "\n")
			sc := strings.Split(text, " ")
			err := dispatch(sc)
			util.HandleError(&err)
		}
	}

}
