package logfiles

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
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

const TimeStampLength = 22

var cpuTemp *regexp.Regexp
var blankTimeStamp string
var gatheredStats []*Series
var cpuTempStats *Series

func init() {
	nullTime = time.Now()
	cpuTemp = regexp.MustCompile("sensord.*temp1: (.*) C")
	blankTimeStamp = strings.Repeat(" ", TimeStampLength)
	gatheredStats = make([]*Series, Processes+1)

}

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
		}
	}
}

func ExtractTimeStamp(line string) (time.Time, error) {
	timstr := line[0:TimeStampLength]
	val, err := time.Parse(time.StampMicro, timstr)
	if err != nil {
		return nullTime, err
	}
	return val, nil
}

func ExtractCPUTemp(line string) (bool, time.Time, float32) {
	if cpuTemp.MatchString(line) {
		at, _ := ExtractTimeStamp(line)
		ts := cpuTemp.FindStringSubmatch(line)
		tv, _ := strconv.ParseFloat(ts[1], 32)
		return true, at, float32(tv)
	}
	return false, nullTime, 0.0
}

func ConnectionStatusLine(line string) (bool, time.Time) {
	if strings.Contains(line, "Connection Status") {
		tv, _ := ExtractTimeStamp(line)
		return true, tv
	}
	return false, nullTime
}

func ConnectionStatusDetailLine(line string) (bool, int, string) {
	if line[0:TimeStampLength] == blankTimeStamp {
		for idx, item := range itemStatusLine {
			pos := strings.Index(line, item)
			if pos > 0 {
				return true, idx, line[pos+len(item) : len(line)]
			}
		}
	}
	return false, 0, ""
}

func ExtractValue(valtype int, val string) float32 {
	if gatheredStats[valtype].Extractor != nil {
		extr := gatheredStats[valtype].Extractor.FindStringSubmatch(val)
		vstr := extr[1]
		val, _ := strconv.ParseFloat(vstr, 32)
		return float32(val)
	}
	return -1.0
}
func AnalyzeFile(rdr io.Reader, base time.Time) {

	var connectionStatus bool = false
	var connectionStatusTime time.Time
	var tempSample Sample

	scanner := bufio.NewScanner(rdr)
	for scanner.Scan() {
		line := scanner.Text()
		if Verbose {
			fmt.Printf("%s\n", line)
		}
		if connectionStatus {
			found, valtype, valstr := ConnectionStatusDetailLine(line)
			if found {
				fmt.Printf("Found detail %d %s\n", valtype, valstr)
				if gatheredStats[valtype] != nil {
					tempSample.At = connectionStatusTime
					//val, _ := strconv.Atoi(strings.TrimSpace(valstr))
					//tempSample.Value = float32(val)
					tempSample.Value = ExtractValue(valtype, strings.TrimSpace(valstr))
					gatheredStats[valtype].Add(tempSample)
					fmt.Printf("Time %v : Type : %s Value %s %f\n", connectionStatusTime, itemStatusLine[valtype], valstr, tempSample.Value)
				} else {
					fmt.Printf("Not gathering the value\n")
				}
			} else {
				connectionStatus = false
			}
		} else {
			found, attime, temp := ExtractCPUTemp(line)
			if found {
				attime := OffsetYear(attime, base)
				if cpuTempStats != nil {
					tempSample.At = attime
					tempSample.Value = temp
					cpuTempStats.Add(tempSample)
				}
				if Verbose {
					fmt.Printf("Time %v : Type : CPU Temp Value %f\n", attime, temp)
				}
			} else {
				found, attime := ConnectionStatusLine(line)
				if found {
					connectionStatus = true
					connectionStatusTime = attime
				}
			}
		}
	}
}
