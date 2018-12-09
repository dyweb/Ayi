// Package Ayi defines interfaces
package Ayi

import (
	"sort"
	"sync"

	"github.com/spf13/cobra"

	"github.com/dyweb/Ayi/config"
	"github.com/dyweb/Ayi/util/logutil"
	"github.com/dyweb/gommon/errors"
)

var log, _ = logutil.NewPackageLoggerAndRegistry()

var (
	appMu        sync.Mutex
	appFactories = make(map[string]AppFactory)
)

// Registry is used to pass configuration and share states between apps
type Registry interface {
	// config
	// ~/.ayi.yml
	HasHomeConfig() bool
	HomeConfig() config.AyiConfig
}

type AppFactory func(Registry) (App, error)

type App interface {
	CobraCommand() *cobra.Command
}

func RegisterAppFactory(name string, factory AppFactory) {
	appMu.Lock()
	defer appMu.Unlock()
	if _, dup := appFactories[name]; dup {
		log.Panicf("RegisterAppFactory is called twice for %s", name)
	}
	appFactories[name] = factory
}

func CreateApp(name string, r Registry) (App, error) {
	appMu.Lock()
	defer appMu.Unlock()
	if f, ok := appFactories[name]; !ok {
		return nil, errors.Errorf("factory %s is not registered", name)
	} else {
		return f(r)
	}
}

func Apps() []string {
	appMu.Lock()
	defer appMu.Unlock()
	list := make([]string, 0, len(appFactories))
	for name := range appFactories {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}
