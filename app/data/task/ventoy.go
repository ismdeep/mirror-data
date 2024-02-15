package task

import (
	"strings"

	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type Ventoy struct {
}

func (receiver *Ventoy) Run() {
	github.FetchReleases("ventoy", "ventoy", "Ventoy", func(s string) bool {
		return strings.Contains(s, "beta")
	})
}
