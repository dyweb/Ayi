package github

import (
	"github.com/dyweb/Ayi"
	"github.com/dyweb/Ayi/util/logutil"
)

const appName = "github"

var log = logutil.NewPackageLogger()

func init() {
	Ayi.RegisterAppFactory(appName, func(r Ayi.Registry) (Ayi.App, error) {
		return NewApp(r)
	})
}
