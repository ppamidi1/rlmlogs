package logfiles

import (
	"fmt"
	"regexp"
	"time"
)

const MIN_SAMPLE_COUNT = 3600 * 8

type Sample struct {
	At    time.Time
	Value float32
}

type Series struct {
	Name      string
	Min       float64
	Max       float64
	Extractor *regexp.Regexp
	Samples   []Sample
}

func New(nm string) *Series {
	var temp = new(Series)
	temp.Name = nm
	//temp.Min = 0
	//temp.Max = 100
	//temp.Samples = make([]Sample, MIN_SAMPLE_COUNT)
	return temp
}

func (ser *Series) Add(s Sample) {
	ser.Samples = append(ser.Samples, s)
}

func (ser *Series) Range() (float64, float64) {
	var minv float32
	var maxv float32
	minv = ser.Samples[0].Value
	maxv = minv
	for _, v := range ser.Samples {
		if v.Value > maxv {
			maxv = v.Value
		}
		if v.Value < minv {
			minv = v.Value
		}
	}
	return float64(minv), float64(maxv)
}

func (ser *Series) SetRange(min float64, max float64) {
	ser.Min = min
	ser.Max = max
}
func (ser *Series) show() {
	fmt.Printf("Series %s Length %d Capacity %d\n", ser.Name, len(ser.Samples), cap(ser.Samples))
	for idx, samp := range ser.Samples {
		fmt.Printf("%000d : %v %f\n", idx, samp.At, samp.Value)
	}
}
