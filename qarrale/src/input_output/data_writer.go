package input_output

import "io"

func WriteData(wd io.Writer, args ...string) error {
	var str string
	for i := 0; i < len(args); i++ {
		str = str + args[i] + " "
	}
	io.WriteString(wd, str+"\n")
	return nil
}
