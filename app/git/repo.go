package git

// Repo is a repository with local path and remote
type Repo struct {
	Remote    Remote
	LocalPath string
}
