package git

import (
	"github.com/dyweb/Ayi"
	"github.com/dyweb/Ayi/util/logutil"
)

const appName = "git"

var log, _ = logutil.NewPackageLoggerAndRegistry()

func init() {
	Ayi.RegisterAppFactory(appName, func(r Ayi.Registry) (Ayi.App, error) {
		return NewApp(r)
	})
}
