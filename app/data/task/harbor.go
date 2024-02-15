package task

import (
	"strings"

	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type Harbor struct {
}

func (receiver *Harbor) Run() {
	github.FetchReleases("harbor", "goharbor", "harbor", func(s string) bool {
		return strings.Contains(s, "-rc") || strings.Contains(s, "-tech-preview")
	})
}
