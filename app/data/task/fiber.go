package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type Fiber struct {
}

func (receiver *Fiber) Run() {
	github.FetchReleases("fiber", "gofiber", "fiber", nil)
}
