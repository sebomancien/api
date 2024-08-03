package logger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	None    LogLevel = 0
	Error   LogLevel = 1
	Warning LogLevel = 2
	Verbose LogLevel = 3
)

var logLevel LogLevel = None

func init() {
	// Read log level from environment variable
	switch os.Getenv("LOG") {
	case "error":
		logLevel = Error
	case "warning":
		logLevel = Warning
	case "verbose":
		logLevel = Verbose
	case "none":
		logLevel = None
	default:
		logLevel = Verbose
	}
}

func LogError(v ...any) {
	if logLevel >= Error {
		log.Printf("[ERROR] %s\n", fmt.Sprint(v...))
	}
}

func LogWarning(v ...any) {
	if logLevel >= Warning {
		log.Printf("[WARNING] %s\n", fmt.Sprint(v...))
	}
}

func LogInfo(v ...any) {
	if logLevel >= Verbose {
		log.Printf("[INFO] %s\n", fmt.Sprint(v...))
	}
}
