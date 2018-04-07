package ayi

import (
	"sync"

	"github.com/spf13/cobra"

	"github.com/dyweb/Ayi/ayi/util/logutil"
	"github.com/dyweb/gommon/errors"
	"sort"
)

var log = logutil.NewPackageLogger()

var (
	appMu        sync.Mutex
	appFactories = make(map[string]AppFactory)
)

type AppFactory func() (App, error)

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

func CreateApp(name string) (App, error) {
	appMu.Lock()
	defer appMu.Unlock()
	if f, ok := appFactories[name]; !ok {
		return nil, errors.Errorf("factory %s is not registered", name)
	} else {
		// TODO: factory may require parameters in the future, i.e. global config instance
		return f()
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
