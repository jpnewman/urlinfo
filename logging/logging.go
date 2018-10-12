// Package logging is a small package for logrus logging.
package logging

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// RootLogger Root Logger
var RootLogger = logrus.New()

// Logger Logrus Entry with ID
var Logger *logrus.Entry

// Log File
var LogFile *os.File

// LogInit Initialize Logrus log
func LogInit(logFile string) {
	id := uuid.New()

	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "info"
	}
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	RootLogger.SetLevel(ll)

	RootLogger.Formatter = &logrus.JSONFormatter{}

	LogFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		RootLogger.Out = LogFile
	} else {
		RootLogger.Info("Failed to create log file, using default stderr")
	}

	fmt.Printf("Log Id: %s\n", id)
	Logger = RootLogger.WithFields(logrus.Fields{
		"id": id,
	})
}

// LogFileClose Log File Close
func LogFileClose() {
	LogFile.Close()
}
