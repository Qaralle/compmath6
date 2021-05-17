package input_output

import (
	"bufio"
	"compmath5/qarrale/src/dto"
	"compmath5/qarrale/src/model"
	"compmath5/qarrale/src/util"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ReadData(rd io.Reader, flag bool) (model.Data, dto.Env, error) {

	reader := bufio.NewReader(rd)

	if flag {

		method := PrepareMethod(reader)
		function := PrepareFunction(reader)

		fmt.Println("Введите интервал [a;b] через пробел: ")

		a, err := reader.ReadString(' ')
		util.HandleError(&err)
		a = strings.TrimSuffix(a, " ")
		aValue, _ := strconv.ParseFloat(a, 64)

		b, err := reader.ReadString('\n')
		util.HandleError(&err)
		b = strings.TrimSuffix(b, "\n")
		bValue, _ := strconv.ParseFloat(b, 64)

		fmt.Println("Введите шаг разбиения h: ")
		h, err := reader.ReadString('\n')
		util.HandleError(&err)
		h = strings.TrimSuffix(h, "\n")
		hValue, _ := strconv.ParseFloat(h, 64)

		fmt.Println("Введите шаг разбиения y0: ")
		y0, err := reader.ReadString('\n')
		util.HandleError(&err)
		y0 = strings.TrimSuffix(y0, "\n")
		y0Value, _ := strconv.ParseFloat(y0, 64)

		return model.Data{Name: "test", A: aValue, B: bValue, Y0: y0Value, H: hValue, Function: function}, dto.Env{Method: method}, nil

	} else {

		method := PrepareMethod(reader)
		function := PrepareFunction(reader)
		return model.Data{Name: "test", A: 1, B: 1, Y0: 1, Function: function}, dto.Env{Method: method}, nil

	}

}
