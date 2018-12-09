package logutil

import (
	ilog "github.com/dyweb/go.ice/ice/util/logutil"
	"github.com/dyweb/gommon/log"
)

const Project = "github.com/dyweb/Ayi"

var logger, registry = log.NewApplicationLoggerAndRegistry(Project)

func Registry() *log.Registry {
	return registry
}

func Logger() *log.Logger {
	return logger
}

func NewPackageLoggerAndRegistry() (*log.Logger, *log.Registry) {
	logger, child := log.NewPackageLoggerAndRegistryWithSkip(Project, 1)
	registry.AddRegistry(child)
	return logger, child
}

// add go.ice log registry as its child
func init() {
	registry.AddRegistry(ilog.Registry())
}
