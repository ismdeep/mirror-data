package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type Ctop struct {
}

func (receiver *Ctop) Run() {
	github.FetchReleases("ctop", "bcicen", "ctop", nil)
}
