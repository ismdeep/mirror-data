package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type GoSec struct {
}

func (receiver *GoSec) Run() {
	github.FetchReleases("gosec", "securego", "gosec", nil)
}
