package util

import (
	"github.com/Sirupsen/logrus"
)

// Logger is the default logger with info level
var Logger = logrus.New()

func init() {
	Logger.Formatter = &logrus.TextFormatter{ForceColors: true}
	Logger.Level = logrus.InfoLevel
}

// UseVerboseLog set logger level to debug
func UseVerboseLog() {
	Logger.Level = logrus.DebugLevel
}
