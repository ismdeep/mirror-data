package main

import "github.com/ismdeep/mirror-data/internal/github"

func main() {
	github.FetchReleases("image-syncer", "AliyunContainerService", "image-syncer")
}
