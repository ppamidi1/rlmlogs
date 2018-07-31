package logfiles

import (
	"fmt"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
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
	//fmt.Printf("Min %f Max %f\n", p.Y.Min, p.Y.Max)
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

func plotSeries(plt *plot.Plot, series *Series, title string) {

	plt.Title.Text = title
	plt.X.Label.Text = "Time (secs)"
	plt.Y.Label.Text = series.Name

	plt.Y.Min, plt.Y.Max = series.Range()
	//fmt.Printf("Min %f Max %f\n", p.Y.Min, p.Y.Max)
	pts := make(plotter.XYs, len(series.Samples))
	for i := range pts {
		pts[i].X = float64(series.Samples[i].At.Sub(series.Samples[0].At) / time.Second)
		pts[i].Y = float64(series.Samples[i].Value)
	}

	err := plotutil.AddLinePoints(plt, series.Name, pts)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func CombinedPlot(series []*Series, fn string, title string) {
	var err error
	plots := make([][]*plot.Plot, numPlots)
	for j := 0; j < numPlots; j++ {
		plots[j] = make([]*plot.Plot, 1)
		plots[j][0], err = plot.New()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}

	img := vgimg.New(7*vg.Inch, 8*vg.Inch)
	dc := draw.New(img)

	t := draw.Tiles{
		Rows:      numPlots,
		Cols:      1,
		PadX:      vg.Millimeter,
		PadY:      vg.Millimeter,
		PadTop:    vg.Points(2),
		PadBottom: vg.Points(2),
		PadLeft:   vg.Points(2),
		PadRight:  vg.Points(2),
	}

	canvases := plot.Align(plots, t, dc)
	pnum := 0
	for _, ser := range series {
		if ser != nil {
			if pnum == 0 {
				plotSeries(plots[pnum][0], ser, title)
			} else {
				plotSeries(plots[pnum][0], ser, "")
			}

			pnum++
		}
	}

	for j := 0; j < numPlots; j++ {
		if plots[j][0] != nil {
			plots[j][0].Draw(canvases[j][0])
		}
	}

	w, err := os.Create(fn + ".png")
	if err != nil {
		panic(err)
	}

	png := vgimg.PngCanvas{Canvas: img}
	if _, err := png.WriteTo(w); err != nil {
		panic(err)
	}
}
