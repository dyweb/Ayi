package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"
	"github.com/dyweb/gommon/util/fsutil"

	"github.com/dyweb/Ayi"
	"github.com/dyweb/Ayi/config"
	"github.com/dyweb/Ayi/util/logutil"

	// apps
	_ "github.com/dyweb/Ayi/app/git"
	_ "github.com/dyweb/Ayi/app/github"
	_ "github.com/dyweb/Ayi/app/web"
	"github.com/dyweb/Ayi/util/configutil"
	"os/user"
	"path/filepath"
)

const (
	myname     = "ayi" // 阿姨
	mydesc     = "Ayi helps you keep things organized"
	configFile = ".ayi.yml"
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
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	cli = icli.New(
		icli.Name(myname),
		icli.Description(mydesc),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	// TODO: go.ice need to disable built in config flag
	r := &registry{}
	apps := Ayi.Apps()
	for _, name := range apps {
		app, err := Ayi.CreateApp(name, r)
		if err != nil {
			return err
		}
		root.AddCommand(app.CobraCommand())
	}
	if err := root.Execute(); err != nil {
		return err
	}
	return nil
}

var _ Ayi.Registry = (*registry)(nil)

type registry struct {
	homeConfigChecked bool
	homeConfig        *config.AyiConfig
}

func (r *registry) HasHomeConfig() bool {
	if !r.homeConfigChecked {
		r.checkHomeConfig()
	}
	return r.homeConfig != nil
}

func (r *registry) HomeConfig() config.AyiConfig {
	if !r.homeConfigChecked {
		r.checkHomeConfig()
	}
	if r.homeConfig != nil {
		return *r.homeConfig
	} else {
		log.Warn("home config is empty")
		return config.AyiConfig{}
	}
}

func (r *registry) checkHomeConfig() {
	r.homeConfigChecked = true
	// TODO: fsutil should have homedir like https://github.com/mitchellh/go-homedir
	u, err := user.Current()
	if err != nil {
		log.Warnf("can't get current user %s", err)
		return
	}
	fp := filepath.Join(u.HomeDir, configFile)
	if !fsutil.FileExists(fp) {
		log.Debugf("config file %s not found", fp)
		return
	}
	cfg := &config.AyiConfig{}
	if err := configutil.LoadYAMLFile(fp, cfg); err != nil {
		log.Warnf("failed to load home config %s: %s", fp, err)
		return
	}
	r.homeConfig = cfg
}
