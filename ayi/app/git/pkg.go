package git

import (
	"github.com/dyweb/Ayi/ayi"
	"github.com/dyweb/Ayi/ayi/util/logutil"
)

const appName = "git"

var log = logutil.NewPackageLogger()

func init() {
	ayi.RegisterAppFactory(appName, func() (ayi.App, error) {
		return NewApp()
	})
}
