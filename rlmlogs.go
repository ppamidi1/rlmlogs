package main

import (
   "./cmd"
)

func main() {
	cmd.Execute()
}

/*
package main

import (
	"fmt"
  "./cmd"
	"logfiles"
  "github.com/spf13/cobra"
)

var verbose bool
var cpuTemp bool
var rtDiag string

var cpuTempSeries *logfiles.Series

func init() {
	flag.BoolVar(&verbose, "-verbose", false, "be verbose")
	flag.BoolVar(&cpuTemp, "-cputemp", false, "plot CPU temperature")
  flag.StringVar(&rtDiag, "-rtdiag" , "" , "output of the rt diag line")
	flag.Parse()

	if verbose {
		fmt.Printf("Will generate following plots:\n")
		if cpuTemp {
			fmt.Printf("\tCPU Temp\n")
		}
	}
	if cpuTemp {
		cpuTempSeries = logfiles.New("CPU Temp. deg C")
	}
}

func main() {

}*/
