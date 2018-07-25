package logfiles

import (
	"fmt"
	"testing"
)

const jsontestline1 string = `Jun 21 03:18:26.397418 RL00122 upload_stats.py[717]: {"vst": 1158601, "itc": 55990, "MemRLC": 64256, "filename": "/var/rlm/dashboard_stats/20180621_0318stats.log", "frc": 55990, "vcc": 52, "ver": 0, "dVST": 20937, "SOM Version": " 1.3.58+rc_v1.3", "MemSystem": 320864256, "fmt": 0, "ts": 1529551315, "Bearer_Change": 1, "machine": "RL00122", "pc": 0, "dPC": 0, "rc": 2, "time_formatted": "2018-06-21 03:18:24", "dRC": 0, "fmc": 0, "rcc": 51, "nmt": 1162200, "Stream_Quality_int": 10, "mmt": 0, "ncc": 440, "Stream_Quality": "STATE_high_quality", "Up Time": 77, "rec": 2, "date": 1529551104, "lasttime": 1529528911, "mmc": 0, "first_from_rlm": "false", "icc": 1188981, "Active Bearer": " Ethernet", "Stream_Quality_change": 1, "dT": 22193, "cpu": 15, "nmc": 0}`

func TestJsonExtract(t *testing.T) {
	Verbose = true
	linedata := JsonExtract(jsontestline1)
	for k, v := range linedata {
		fmt.Printf("Key: %s Type: ", k)
		switch vtype := v.(type) {
		case bool:
			fmt.Printf("bool ")
		case float64:
			fmt.Printf("float64 ")
		case string:
			fmt.Printf("string ")
		case nil:
			fmt.Printf("nil ")
		case []interface{}:
			fmt.Printf("array ")
		default:
			fmt.Printf("Cannot decode %v ", vtype)
		}
		fmt.Printf("Value: %v\n", v)
	}
}
