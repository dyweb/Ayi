package git

import (
	"github.com/dyweb/Ayi"
	"github.com/dyweb/Ayi/util/logutil"
)

const appName = "git"

var log = logutil.NewPackageLogger()

func init() {
	Ayi.RegisterAppFactory(appName, func() (Ayi.App, error) {
		return NewApp()
	})
}
