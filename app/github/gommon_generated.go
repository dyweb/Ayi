// Code generated by gommon from app/github/gommon.yml DO NOT EDIT.

package github

import dlog "github.com/dyweb/gommon/log"

func (a *App) SetLogger(logger *dlog.Logger) {
	a.log = logger
}

func (a *App) GetLogger() *dlog.Logger {
	return a.log
}

func (a *App) LoggerIdentity(justCallMe func() dlog.Identity) dlog.Identity {
	return justCallMe()
}
