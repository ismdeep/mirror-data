package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type Pandoc struct {
}

func (receiver *Pandoc) Run() {
	github.FetchReleases("pandoc", "jgm", "pandoc", nil)
}
