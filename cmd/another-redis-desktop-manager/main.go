package main

import "github.com/ismdeep/mirror-data/internal/github"

func main() {
	github.FetchReleases("another-redis-desktop-manager", "qishibo", "AnotherRedisDesktopManager")
}
