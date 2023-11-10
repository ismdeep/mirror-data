package ardm

import "github.com/ismdeep/mirror-data/internal/github"

func Run() error {
	github.FetchReleases("another-redis-desktop-manager", "qishibo", "AnotherRedisDesktopManager")
	return nil
}
