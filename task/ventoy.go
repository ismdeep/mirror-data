package task

import "github.com/ismdeep/mirror-data/internal/github"

type Ventoy struct {
}

func (receiver *Ventoy) Run() {
	github.FetchReleases("ventoy", "ventoy", "Ventoy")
}
