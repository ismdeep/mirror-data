package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type Rclone struct {
}

func (receiver *Rclone) Run() {
	github.FetchReleases("rclone", "rclone", "rclone", nil)
}
