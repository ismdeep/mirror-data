package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type GolangciLint struct {
}

func (receiver *GolangciLint) Run() {
	github.FetchReleases("golangci-lint", "golangci", "golangci-lint", nil)
}
