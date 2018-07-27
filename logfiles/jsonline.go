package logfiles

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var taggedStats map[string]*Series
var jsonTags = [...]string{
	"vst",
	"itc",
	"MemRLC",
	"filename",
	"frc",
	"vcc",
	"ver",
	"dVST",
	"SOM Version",
	"MemSystem",
	"fmt",
	"ts",
	"Bearer_Change",
	"machine",
	"pc",
	"dPC",
	"rc",
	"time_formatted",
	"dRC",
	"fmc",
	"rcc",
	"nmt",
	"Stream_Quality_int",
	"mmt",
	"ncc",
	"Stream_Quality",
	"Up Time",
	"rec",
	"date",
	"lasttime",
	"mmc",
	"first_from_rlm",
	"icc",
	"Active Bearer",
	"Stream_Quality_change",
	"dT",
	"cpu",
	"nmc"}

const keyTag string = "SOM Version"

/*
upload_stats.py[717]: {"vst": 1158601, "itc": 55990, "MemRLC": 64256, "filename": "/var/rlm/dashboard_stats/20180621_0318stats.log", "frc": 55990, "vcc": 52, "ver": 0, "dVST": 20937, "SOM Version": " 1.3.58+rc_v1.3", "MemSystem": 320864256, "fmt": 0, "ts": 1529551315, "Bearer_Change": 1, "machine": "RL00122", "pc": 0, "dPC": 0, "rc": 2, "time_formatted": "2018-06-21 03:18:24", "dRC": 0, "fmc": 0, "rcc": 51, "nmt": 1162200, "Stream_Quality_int": 10, "mmt": 0, "ncc": 440, "Stream_Quality": "STATE_high_quality", "Up Time": 77, "rec": 2, "date": 1529551104, "lasttime": 1529528911, "mmc": 0, "first_from_rlm": "false", "icc": 1188981, "Active Bearer": " Ethernet", "Stream_Quality_change": 1, "dT": 22193, "cpu": 15, "nmc": 0}

{
	"vst": 1158601,
	"itc": 55990,
	"MemRLC": 64256,
	"filename": "/var/rlm/dashboard_stats/20180621_0318stats.log",
	"frc": 55990,
	"vcc": 52,
	"ver": 0,
	"dVST": 20937,
	"SOM Version": " 1.3.58+rc_v1.3",
	"MemSystem": 320864256,
	"fmt": 0,
	"ts": 1529551315,
	"Bearer_Change": 1,
	"machine": "RL00122",
	"pc": 0,
	"dPC": 0,
	"rc": 2,
	"time_formatted": "2018-06-21 03:18:24",
	"dRC": 0,
	"fmc": 0,
	"rcc": 51,
	"nmt": 1162200,
	"Stream_Quality_int": 10,
	"mmt": 0,
	"ncc": 440,
	"Stream_Quality": "STATE_high_quality",
	"Up Time": 77,
	"rec": 2,
	"date": 1529551104,
	"lasttime": 1529528911,
	"mmc": 0,
	"first_from_rlm": "false",
	"icc": 1188981,
	"Active Bearer": " Ethernet",
	"Stream_Quality_change": 1,
	"dT": 22193,
	"cpu": 15,
	"nmc": 0
}
*/

var jsonlineExp *regexp.Regexp

func init() {
	//jsonlineExp = regexp.MustCompile(".*upload_stats\\.py\\.*\\:\\.*(\\{.*\\})")
	jsonlineExp = regexp.MustCompile("(\\{.*\\})")
	taggedStats = make(map[string]*Series)
}

func IsJsonStatLine(full string) (bool, time.Time) {
	if strings.Contains(full, keyTag) {
		ts, _ := ExtractTimeStamp(full)
		return true, ts
	}
	return false, nullTime
}

func JsonExtract(line string) map[string]interface{} {
	jsoncomps := jsonlineExp.FindStringSubmatch(line)
	var split interface{}
	json.Unmarshal([]byte(jsoncomps[0]), &split)
	return split.(map[string]interface{})
}

func AnalyzeJsonLine(full string) {
	yes, ts := IsJsonStatLine(full)
	if !yes {
		return
	}
	CollectStatsJsonLine(ts, full)
}

func CollectStatsJsonLine(t time.Time, line string) {
	var samp Sample
	values := JsonExtract(line)
	for k, v := range values {
		if taggedStats[k] != nil {
			val, ok := v.(float64)
			if ok {
				samp.At = t
				samp.Value = float32(val)
				taggedStats[k].Add(samp)
			}
		}
	}
}

func SetupTag(tag string) {
	taggedStats[tag] = New(tag)
}

func ValidTag(t string) bool {
	for _, tag := range jsonTags {
		if tag == t {
			return true
		}
	}
	return false
}

func ShowValidTags() {
	for _, tag := range jsonTags {
		fmt.Printf("%s\n", tag)
	}
}
