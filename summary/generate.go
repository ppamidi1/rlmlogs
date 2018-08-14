package summary

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

var Verbose bool

var logFileNames *regexp.Regexp
var rlmsWithNoLogs []string

func init() {
	logFileNames = regexp.MustCompile("[0-9]+_[0-9]+.log")
}

func PrintSummaries() {
	fmt.Printf("RLMs which had no logs\n")
	for _, rlm := range rlmsWithNoLogs {
		fmt.Printf("%s\n", rlm)
	}
}
func SummarizeRLM(rlmdir string) {
	if Verbose {
		fmt.Printf("Summarizing %s\n", rlmdir)
	}
	jdir := path.Join(rlmdir, "journal")
	var numlogfiles int
	jfs, err := ioutil.ReadDir(jdir)
	if err != nil {
		fmt.Printf("%s - %s\n", rlmdir, err.Error())
		return
	}
	for _, jfs := range jfs {
		if logFileNames.MatchString(jfs.Name()) {
			if Verbose {
				fmt.Printf("Found a log file %s\n", jfs.Name())
			}
			numlogfiles++
		}
	}
	if numlogfiles == 0 {
		rlmsWithNoLogs = append(rlmsWithNoLogs, path.Base(rlmdir))
		if Verbose {
			fmt.Printf("%s did not have any logfiles\n", path.Base(rlmdir))
		}
	}
}
func Generate(toplevel string) {
	if Verbose {
		fmt.Printf("Generating summary of RLMs rooted at %s\n", toplevel)
	}
	st, err := os.Stat(toplevel)
	if err == nil && st.IsDir() {
		rlms, err := ioutil.ReadDir(toplevel)
		if err != nil {
			fmt.Printf("%s\n", err.Error)
			return
		}
		for _, rlm := range rlms {
			if rlm.IsDir() {
				SummarizeRLM(path.Join(toplevel, rlm.Name()))
			} else {
				if Verbose {
					fmt.Printf("Skipping %s. Not a directory\n", rlm.Name())
				}
			}
		}
		PrintSummaries()
	} else {
		fmt.Printf("%s is not a directory\n", toplevel)
	}
}
