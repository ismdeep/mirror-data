package task

import (
	"strings"

	"github.com/ismdeep/mirror-data/internal/github"
)

type GitForWindows struct {
}

func (receiver *GitForWindows) Run() {
	github.FetchReleases("git-for-windows", "git-for-windows", "git", func(s string) bool {
		return strings.Contains(s, "-rc") || strings.Contains(s, "prerelease-")
	})
}
