package logfiles

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var Verbose bool = false

var zippedLogs regexp.Regexp
var journalLogs regexp.Regexp

func init() {

}

func analyzeZipFile(fn string) {
	var filenames []string
	zipdate, err := DateOfZippedLogs(fn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Analyzing zip file of %s\n", zipdate.Format("Jan 2 2006"))
	zipdate = zipdate.AddDate(0, 0, -1)
	reader, err := zip.OpenReader(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	for _, file := range reader.File {
		if Verbose {
			fmt.Printf("%30s Size %10d Compressed %10d CRC %10x\n",
				file.Name,
				file.UncompressedSize,
				file.CompressedSize,
				file.CRC32)
		}
		if strings.Contains(file.Name, "stats.log") {
			if Verbose {
				log.Printf("Skipping %s\n", file.Name)
			}
		} else {
			filenames = append(filenames, file.Name)
		}
	}
	sort.Strings(filenames)

	for _, fn := range filenames {
		for _, zfile := range reader.File {
			if zfile.Name == fn {
				zrdr, _ := zfile.Open()
				defer zrdr.Close()
				fdate, _ := DateOfLogfile(zfile.Name)
				AnalyzeFile(zrdr, fdate)
			}
		}
	}
}

func analyzeLogFile(fn string) {
	logdate, err := DateOfLogfile(fn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Analyzing log file of %s\n", logdate.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	lf, err := os.Open(fn)
	defer lf.Close()
	if err == nil {
		AnalyzeFile(lf, logdate)
	}
}

func Analyze(nm string) {
	ext := filepath.Ext(nm)
	switch ext {
	case ".zip":
		analyzeZipFile(nm)
	case ".log":
		analyzeLogFile(nm)
	default:
		log.Printf("%s is not a type of logfile I can analyze\n", nm)
	}
}

func GeneratePlots(nm string, title string) {
	if cpuTempStats != nil {
		cpuTempStats.Plot("cputemp"+nm, title)
	}
	for idx, chart := range gatheredStats {
		if chart != nil {
			chart.Plot(itemNames[idx]+nm, title)
		}
	}
}
