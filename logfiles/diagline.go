package logfiles

import (
	"fmt"
	"regexp"
)

const (
	AM = iota
	AS
	CS
	SS
	VS
	AB
	IP
	FI
	FO
	FD
	CPU_Usage
	Up_Time
	Pct_Mem_Used_sys
	Mem_Used_sys
	Mem_Free_sys
	Mem_Used_rlc
	Processes
)

var itemNames = [...]string{
	"AM",
	"AS",
	"CS",
	"SS",
	"VS",
	"AB",
	"IP",
	"FI",
	"FO",
	"FD",
	"CPU_Usage",
	"Up_Time",
	"Pct_Mem_Used_sys",
	"Mem_Used_sys",
	"Mem_Free_sys",
	"Mem_Used_rlc",
	"Processes"}

var itemStatusLine = [...]string{
	"AM=",
	"AS=",
	"CS=",
	"SS=",
	"VS=",
	"AB=",
	"IP=",
	"FI=",
	"FO=",
	"FD=",
	"CPU Usage:",
	"Up Time:",
	"% Mem Used(sys):",
	"Mem Used(sys):",
	"Mem Free(sys):",
	"Mem Used(rlc):",
	"Processes:"}

/*
	Jun 20 23:29:39.289159 RL00122 rlc[694]: RLM RL00122 Connection Status
	                                         AM=2
	                                         AS=STATE_started
	                                         CS=STATE_connected
	                                         SS=STATE_established
	                                         VS=STATE_high_quality
	                                         AB=Ethernet
	                                         IP=192.168.128.37
	                                         FI=146308
	                                         FO=146302
	                                         FD=6
	                                         CPU Usage: 55%
	                                         Up Time: 14496
	                                         % Mem Used(sys):39.07781
	                                         Mem Used(sys):  411009024
	                                         Mem Free(sys):  640761856
	                                         Mem Used(rlc):  48696
	                                         Processes: 130  */
var numPlots int

func ValidItem(nm string) bool {
	for _, opt := range itemNames {
		if nm == opt {
			return true
		}
	}
	return false
}
func ShowValidItems() {
	for _, opt := range itemNames {
		fmt.Printf("%s\n", opt)
	}
}
func Index(nm string) int {
	for i, opt := range itemNames {
		if nm == opt {
			return i
		}
	}
	return -1
}

func SetupStats(nm string) {
	if ValidItem(nm) {
		numPlots++
		idx := Index(nm)
		gatheredStats[idx] = New(nm)
		switch idx {
		case FI:
			gatheredStats[FI].Extractor = regexp.MustCompile("([0-9]+)")
		case FO:
			gatheredStats[FO].Extractor = regexp.MustCompile("([0-9]+)")
		case FD:
			gatheredStats[FD].Extractor = regexp.MustCompile("([0-9]+)")
		case CPU_Usage:
			gatheredStats[CPU_Usage].Extractor = regexp.MustCompile("([0-9]+)%")
		case Up_Time:
			gatheredStats[Up_Time].Extractor = regexp.MustCompile("([0-9]+)")
		case Pct_Mem_Used_sys:
			gatheredStats[Pct_Mem_Used_sys].Extractor = regexp.MustCompile("([0-9]+\\.[0-9]+)")
		case Mem_Used_sys:
			gatheredStats[Mem_Used_sys].Extractor = regexp.MustCompile("([0-9]+)")
		case Mem_Free_sys:
			gatheredStats[Mem_Free_sys].Extractor = regexp.MustCompile("([0-9]+)")
		case Mem_Used_rlc:
			gatheredStats[Mem_Used_rlc].Extractor = regexp.MustCompile("([0-9]+)")
		case Processes:
			gatheredStats[Processes].Extractor = regexp.MustCompile("([0-9]+)")
		}
	} else {
		if nm == "CPUTemp" {
			cpuTempStats = New("CPUTemp")
			numPlots++
		}
	}
}
