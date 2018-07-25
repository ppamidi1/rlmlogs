package cmd

import (
	"fmt"
	"strings"
	"time"

	"../logfiles"

	"github.com/spf13/cobra"
)

var PlotFileName string
var RLMId string
var RtDiagVar string
var JsonDiagVar string
var RtDiagOptions []string
var JsonDiagOptions []string
var OptCPUTemp bool
var longHelp string

var plotCmd = &cobra.Command{
	Use:   "plot",
	Short: "generate a plot of the specified log item",
	Long:  longHelp,
	Run:   execPlotCmd,
}

func init() {

	longHelp = `This is the long help`
	plotCmd.Flags().StringVarP(&JsonDiagVar, "jsondiag", "j", "", "extract dashboard stats from the JSON string uploads")
	plotCmd.Flags().StringVarP(&RtDiagVar, "rtdiag", "r", "", "plot rt rlm_diag_data. provide list")
	plotCmd.Flags().BoolVarP(&OptCPUTemp, "cputemp", "t", false, "plot CPU temperature")
	plotCmd.Flags().StringVarP(&RLMId, "rlmid", "i", "RL00000", "Id of the RLM")
	plotCmd.Flags().StringVarP(&PlotFileName, "plotfile", "p", "plot.png", "Output file name")

	rootCmd.AddCommand(plotCmd)
}

func execPlotCmd(cmd *cobra.Command, args []string) {
	if Verbose {
		fmt.Printf("Option CPUTemp: %v\n", OptCPUTemp)
	}
	if OptCPUTemp {
		logfiles.SetupStats("CPUTemp")
	}
	if Verbose {
		fmt.Printf("Diag Variables: %s\n", RtDiagVar)
	}
	if len(RtDiagVar) > 0 {
		RtDiagOptions = strings.Split(RtDiagVar, ",")
		for _, opt := range RtDiagOptions {
			if Verbose {
				fmt.Printf("Analyzing %s\n", opt)
			}
			if logfiles.ValidItem(opt) {
				fmt.Printf("\t%s\n", opt)
				logfiles.SetupStats(opt)
			} else {
				fmt.Printf("\t%s is not a valid rtdiag item\n", opt)
				logfiles.ShowValidItems()
				return
			}
		}
	}
	if len(JsonDiagOptions) > 0 {
		JsonDiagOptions = strings.Split(JsonDiagVar, ",")
	}

	if len(args) < 1 {
		fmt.Printf("Need some files to process\n")
		return
	}

	for _, arg := range args {
		fmt.Printf("Processing %s\n", arg)
		logfiles.Analyze(arg)
	}
	basedate, _ := logfiles.DateOfZippedLogs(args[0])
	logfiles.GeneratePlots(PlotFileName, "RLM "+RLMId+" "+basedate.Format(time.ANSIC))
}
