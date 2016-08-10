package git

import (
	"github.com/dyweb/Ayi/util"
)

// CloneFromURL will clone a repository based on short url or normal git clone url
func CloneFromURL(repoURL string) (Repo, error) {
	// get the remote from url
	r, err := NewFromURL(repoURL)
	log.Debug(r)
	// TODO: may wrap the error to have stack included
	if err != nil {
		return Repo{}, err
	}
	return Clone(r)
}

// Clone clones a remote git repo
func Clone(r Remote) (Repo, error) {
	// log.Info("git clone " + r.GetSSH() + " " + GetCloneDirectory(r))
	repo := Repo{Remote: r, LocalPath: GetCloneDirectory(r)}
	cmdStr := "git clone " + r.GetSSH() + " " + repo.LocalPath
	log.Info(cmdStr)
	err := util.RunCommand(cmdStr)
	if err != nil {
		// TODO: may wrap the error
		return repo, err
	}
	return repo, nil
}
