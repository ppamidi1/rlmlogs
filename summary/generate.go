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

type LogDetails struct {
	Name string
	NumberOfLogFiles int
}
var rlmsWithLogs []LogDetails

func init() {
	logFileNames = regexp.MustCompile("\d{8}_\d{4}.log")
	rlmDirNames = regexp.MustCompile("(?:.+\/)?RL\d{5}")
}

func PrintSummaries() {
	fmt.Printf("RLMs reporting with logs\n")
	for _, rlm := range rlmsWithLogs {
		fmt.Printf("%s    - %d\n", rlm.Name, rlm.NumberOfLogFiles)
	}

	fmt.Printf("RLMs which had no logs\n")
	for _, rlm := range rlmsWithNoLogs {
		fmt.Printf("%s\n", rlm)
	}
}
func SummarizeRLM(rlmdir string) {
	if !rlmDirNames.MatchString(rlmdir) {
		if Verbose {
			fmt.Printf("%s not an RLM directory\n", rlmdir)
		}
		return
	}
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
	} else {
		var temp LogDetails
		temp.Name = path.Base(rlmdir)
		temp.NumberOfLogFiles = numlogfiles
		rlmsWithLogs = append(rlmsWithLogs , temp)
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
