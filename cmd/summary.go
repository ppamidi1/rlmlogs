package cmd

import (
	"fmt"

	"../summary"
	"github.com/spf13/cobra"
)

var TopLevelDir string

func init() {
	summaryCmd.Flags().StringVarP(&TopLevelDir, "top", "t", "", "toplevel dir where RLM data is found")
	rootCmd.AddCommand(summaryCmd)
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summary of RLMs",
	Long:  `Summary of which RLMs are sending journals, etc`,
	Run:   execSummaryCmd,
}

func execSummaryCmd(cmd *cobra.Command, args []string) {
	summary.Verbose = Verbose
	if len(TopLevelDir) == 0 {
		fmt.Println("Please provide the toplevel dir where RLM data is recorded")
		return
	}
	summary.Generate(TopLevelDir)
}
