package task

import "github.com/ismdeep/mirror-data/internal/github"

type GitForWindows struct {
}

func (receiver *GitForWindows) Run() {
	github.FetchReleases("git-for-windows", "git-for-windows", "git")
}
