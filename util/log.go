package util

import (
	"github.com/uber-go/zap"
)

// Logger is the default logger with info level
var Logger = zap.NewJSON()

func init() {
	Logger.SetLevel(zap.InfoLevel)
}

// UseVerboseLog set logger level to debug
func UseVerboseLog() {
	Logger.SetLevel(zap.DebugLevel)
}
