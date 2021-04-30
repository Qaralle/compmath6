package input_output

import (
	"bufio"
	model2 "compmath4/qarrale/src/model"
	"compmath4/qarrale/src/util"
	"io"
	"strconv"
	"strings"
)

func ReadData(rd io.Reader) (model2.Data, error) {

	reader := bufio.NewReader(rd)
	text, err := reader.ReadString('\n')
	util.HandleError(&err)
	text = strings.TrimSuffix(text, "\n")

	n, _ := strconv.Atoi(text)

	text, err = reader.ReadString('\n')
	util.HandleError(&err)
	text = strings.TrimSuffix(text, "\n")

	x, _ := strconv.ParseFloat(text, 64)

	nodes_x := make([]float64, n)
	nodes_y := make([]float64, n)
	for i := 0; i < n; i++ {
		text_x, err := reader.ReadString(' ')
		util.HandleError(&err)
		text_x = strings.TrimSuffix(text_x, " ")
		nodes_x[i], _ = strconv.ParseFloat(text_x, 32)
		text_y, err1 := reader.ReadString('\n')
		util.HandleError(&err1)
		text_y = strings.TrimSuffix(text_y, "\n")
		nodes_y[i], _ = strconv.ParseFloat(text_y, 32)
	}

	return model2.Data{"test", n, nodes_x, nodes_y, 0, x}, nil

}
