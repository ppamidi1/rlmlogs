package logfiles

import (
	"fmt"
	"testing"
	"time"
)

func TestDateId(t *testing.T) {
	t1 := time.Date(1986, time.April, 1, 0, 0, 0, 0, time.UTC)
	if DateId(t1) == "19860401" {
		t.Log("basic tests passed")
	} else {
		t.Error("basic test failure")
	}
	t.Logf("Today Id %s should be %s\n", DateId(time.Now()), TodayId())
	t.Logf("Yesterday %s\n", YesterdayId())
}

func TestParseId(t *testing.T) {
	str := "19860401"
	t1, err := ParseId(str)
	if err != nil {
		t.Logf("Error %s parsing %s\n", err.Error(), str)
	}
	t.Logf("%s results in %s\n", str, DateId(t1))
	if t1 == time.Date(1986, time.April, 1, 0, 0, 0, 0, time.UTC) {
		t.Log("basic tests passed")
	} else {
		t.Error("basic parse tests failed")
	}
}

func TestDateOfLogfile(t *testing.T) {
	ts, err := DateOfLogfile("20180621_0208.log")
	if err == nil {
		t.Log(ts)
	} else {
		t.Log(err)
	}
	ts, err = DateOfLogfile("20180621.log")
	if err == nil {
		t.Log(ts)
	} else {
		t.Log(err)
	}
}

func TestDateOfZippedLogs(t *testing.T) {
	ts, err := DateOfZippedLogs("20180621.zip")
	if err == nil {
		fmt.Println(ts)
	} else {
		fmt.Println(err)
	}
	ts, err = DateOfLogfile("20180621.log")
	if err == nil {
		fmt.Println(ts)
	} else {
		fmt.Println(err)
	}
}
