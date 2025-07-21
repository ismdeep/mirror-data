package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type Dbeaver struct {
}

func (receiver *Dbeaver) Run() {
	github.FetchReleases("dbeaver", "dbeaver", "dbeaver", nil)
}
