package task

import "github.com/ismdeep/mirror-data/internal/github"

type Harbor struct {
}

func (receiver *Harbor) Run() {
	github.FetchReleases("harbor", "goharbor", "harbor")
}
