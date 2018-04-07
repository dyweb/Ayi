package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"

	"github.com/dyweb/Ayi/ayi"
	_ "github.com/dyweb/Ayi/ayi/app/git"
	"github.com/dyweb/Ayi/ayi/util/logutil"
)

const (
	myname = "ayi" // 阿姨
)

var log = logutil.Registry

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}
var cli *icli.Root

func main() {
	cli = icli.New(
		icli.Name(myname),
		icli.Description("Ayi helps you keep things organized"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	apps := ayi.Apps()
	for _, name := range apps {
		app, err := ayi.CreateApp(name)
		if err != nil {
			Err(err)
		}
		root.AddCommand(app.CobraCommand())
	}
	if err := root.Execute(); err != nil {
		Err(err)
	}
}

func Err(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
