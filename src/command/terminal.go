package command

import (
	"bufio"
	"compmath4/src/util"
	"os"
	"strings"
)

func Start(){
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		util.HandleError(&err)
		if text != "" {
			text=strings.TrimSuffix(text,"\n")
			sc:=strings.Split(text," ")
			err := dispatch(sc)
			util.HandleError(&err)
		}
	}




}