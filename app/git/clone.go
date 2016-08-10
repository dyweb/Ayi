package git

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
	log.Info(r.GetSSH())
	log.Info("git clone " + r.GetSSH() + " " + GetRepoBasePath() + "/" + r.Repo)
	return nil
}
