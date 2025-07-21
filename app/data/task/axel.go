package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type Axel struct {
}

func (receiver *Axel) Run() {
	github.FetchReleases("axel", "axel-download-accelerator", "axel", nil)
}
