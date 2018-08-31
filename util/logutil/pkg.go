package logutil

import (
	ilog "github.com/dyweb/go.ice/ice/util/logutil"
	"github.com/dyweb/gommon/log"
)

// Registry is the root logger of Ayi application
var Registry = log.NewApplicationLogger()

// NewPackageLogger create a new logger for the calling package using Registry as its parent
func NewPackageLogger() *log.Logger {
	l := log.NewPackageLoggerWithSkip(1)
	Registry.AddChild(l)
	return l
}

// add go.ice log registry as its child
func init() {
	Registry.AddChild(ilog.Registry)
}
