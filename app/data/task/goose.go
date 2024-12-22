package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type Goose struct {
}

func (receiver *Goose) Run() {
	github.FetchReleases("goose", "pressly", "goose", nil)
}
