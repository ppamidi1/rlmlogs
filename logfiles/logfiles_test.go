package logfiles

import (
	"fmt"
	"testing"
)

func TestAnalyze(t *testing.T) {
	Verbose = true
	//	Analyze("20180701.zip")
	//	Analyze("20180822_1022.log")
	//	Analyze("badname.txt")
}

func TestAnalyzeMemUsedRlc(t *testing.T) {
	Verbose = true
	SetupStats("Mem_Used_rlc")
	idx := Index("Mem_Used_rlc")
	if idx >= 0 {
		fmt.Printf("Mem_Used_rlc idx %d\n", idx)
		if gatheredStats[idx] == nil {
			fmt.Printf("Stats array not setup\n")
		} else {
			fmt.Printf("Stats array setup successful\n")
		}
	} else {
		fmt.Printf("Bad index for Mem_Used_rlc\n")
	}
	//Analyze("20180822_1022.log")
	//GeneratePlots("mem_used", "Testing Mem_Used_rlc")
}
