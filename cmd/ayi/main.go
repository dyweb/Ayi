package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"

	"github.com/dyweb/Ayi"
	_ "github.com/dyweb/Ayi/app/git"
	"github.com/dyweb/Ayi/util/logutil"
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
	apps := Ayi.Apps()
	for _, name := range apps {
		app, err := Ayi.CreateApp(name)
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
