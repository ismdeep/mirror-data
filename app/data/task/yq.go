package task

import "github.com/ismdeep/mirror-data/app/data/internal/github"

type Yq struct {
}

func (receiver *Yq) Run() {
	github.FetchReleases("yq", "mikefarah", "yq", nil)
}
