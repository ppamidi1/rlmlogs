package logfiles

import (
	"fmt"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func (ser *Series) Plot(fn string, title string) {

	p, err := plot.New()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	p.Title.Text = title
	p.X.Label.Text = "Time (secs)"
	p.Y.Label.Text = ser.Name

	p.Y.Min, p.Y.Max = ser.Range()
	fmt.Printf("Min %f Max %f\n", p.Y.Min, p.Y.Max)
	pts := make(plotter.XYs, len(ser.Samples))
	for i := range pts {
		pts[i].X = float64(ser.Samples[i].At.Sub(ser.Samples[0].At) / time.Second)
		pts[i].Y = float64(ser.Samples[i].Value)
	}

	err = plotutil.AddLinePoints(p, ser.Name, pts)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 4*vg.Inch, fn); err != nil {
		panic(err)
	}
}
