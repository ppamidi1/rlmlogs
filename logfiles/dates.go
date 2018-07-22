package logfiles

import (
	"errors"
	"fmt"
	"log"
	"path"
	"regexp"
	"strconv"
	"time"
)

var nullTime time.Time

var logfileName *regexp.Regexp
var zippedlogs *regexp.Regexp

func init() {
	logfileName = regexp.MustCompile("([0-9]{8})_[0-9]{4}.log")
	zippedlogs = regexp.MustCompile("([0-9]*).zip")
}

func DateId(t time.Time) string {
	yy := t.Year()
	mm := t.Month()
	dd := t.Day()
	return fmt.Sprintf("%4d", yy) + fmt.Sprintf("%02d", mm) + fmt.Sprintf("%02d", dd)
}

func ParseId(id string) (time.Time, error) {
	if len(id) < 8 {
		return time.Now(), errors.New("Date ID string too short")
	}
	yys, _ := strconv.Atoi(id[0:4])
	mms, _ := strconv.Atoi(id[4:6])
	dds, _ := strconv.Atoi(id[6:8])
	if Verbose {
		fmt.Printf("%d %d %d\n", yys, mms, dds)
	}
	return time.Date(yys, time.Month(mms), dds, 0, 0, 0, 0, time.UTC), nil
}
func TodayId() string {
	return DateId(time.Now())
}

func YesterdayId() string {
	return PreviousDateId(time.Now())
}

func PreviousDateId(t time.Time) string {
	return DateId(t.AddDate(0, 0, -1))
}

func DateOfLogfile(fn string) (time.Time, error) {
	bn := path.Base(fn)
	if logfileName.MatchString(bn) {
		log.Printf("File name %s is in the right format\n", bn)
		yyyymmdd := logfileName.FindStringSubmatch(bn)
		fmt.Printf("Date of log %s\n", yyyymmdd[1])
		zdt, err := time.Parse("20060102", yyyymmdd[1])
		return zdt, err
	}
	fmt.Printf("File name %s is not named as expected\n", bn)
	return nullTime, errors.New("File name syntax incorrect")
}

func DateOfZippedLogs(fn string) (time.Time, error) {
	bn := path.Base(fn)
	if zippedlogs.MatchString(bn) {
		if Verbose {
			fmt.Printf("File name %s is in the right format\n", bn)
		}
		yyyymmdd := zippedlogs.FindStringSubmatch(bn)
		fmt.Printf("Date of log %s\n", yyyymmdd[1])
		zdt, err := time.Parse("20060102", yyyymmdd[1])
		return zdt, err
	}
	fmt.Printf("File name %s is not named as expected\n", bn)
	return nullTime, errors.New("File name syntax incorrect")
}

func OffsetYear(tim time.Time, ot time.Time) time.Time {
	return time.Date(ot.Year(), time.Month(tim.Month()), tim.Day(), tim.Hour(), tim.Minute(), tim.Second(), tim.Nanosecond(), tim.Location())
}
