package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type AnotherRedisDesktopManager struct {
}

func (receiver *AnotherRedisDesktopManager) Run() {
	github.FetchReleases("another-redis-desktop-manager", "qishibo", "AnotherRedisDesktopManager", nil)
}
