package graphic

import (
	"compmath5/qarrale/src/function"
	"compmath5/qarrale/src/model"
	"golang.org/x/exp/rand"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func DrawFunction(function function.Function, a, b float64) *plot.Plot {
	p := plot.New()

	p.Title.Text = "Functions"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	quad := plotter.NewFunction(function.F_clear)
	quad.Color = color.RGBA{R: uint8(rand.Intn(255-0) + 0), B: uint8(rand.Intn(255-0) + 0), A: uint8(rand.Intn(255-0) + 0)}

	p.Add(quad)

	p.X.Min = a
	p.X.Max = b

	return p
}

func Draw(data model.Data, p *plot.Plot) {
	pts := make(plotter.XYs, len(data.XValues))

	for i := 0; i < len(data.XValues); i++ {
		pts[i].X = data.XValues[i]
		pts[i].Y = data.YValues[i]
	}

	_ = plotutil.AddLinePoints(p, "dots", pts)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "functions.png"); err != nil {
		panic(err)
	}
}
