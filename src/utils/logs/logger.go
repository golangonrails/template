/*
  Logger definition
*/
package logs

import (
	"log"
	"os"
)

var DefaultWriter = os.Stdout

func init() {
	log.SetFlags(log.Flags() | log.Lmicroseconds | log.Lshortfile)
}

func LogToFile(logFile string) {
	if outfile, err := os.OpenFile(logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666); err != nil {
		log.Fatalf("Open LogFile '%v' Failed: %v\n", logFile, err.Error())
	} else {
		DefaultWriter = outfile
	}
}

func Logger(prefix string) func() *log.Logger {
	return func() *log.Logger {
		return log.New(DefaultWriter, prefix, log.Flags())
	}
}

// V Verbose
var V = Logger("[Verbose] ")

// W Warning
var W = Logger("[Warning] ")

// E Error
var E = Logger("[Error] ")

// I Info
var I = Logger("[Info] ")

// D Debug
var D = Logger("[Debug] ")
