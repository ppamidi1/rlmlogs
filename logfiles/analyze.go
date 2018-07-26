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
