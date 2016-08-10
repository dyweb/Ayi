package git

import (
	"github.com/dyweb/Ayi/util"
)

// CloneFromURL will clone a repository based on short url or normal git clone url
func CloneFromURL(repoURL string) error {
	// get the remote from url
	r, err := NewFromURL(repoURL)
	log.Debug(r)
	// TODO: may wrap the error to have stack included
	if err != nil {
		return err
	}
	return Clone(r)
}

// Clone clones a remote git repo
func Clone(r Remote) error {
	// log.Info("git clone " + r.GetSSH() + " " + GetCloneDirectory(r))
	cmdStr := "git clone " + r.GetSSH() + " " + GetCloneDirectory(r)
	log.Info(cmdStr)
	return util.RunCommand(cmdStr)
}
