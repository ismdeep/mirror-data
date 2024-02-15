package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type ImageSyncer struct {
}

func (receiver *ImageSyncer) Run() {
	github.FetchReleases("image-syncer", "AliyunContainerService", "image-syncer", nil)
}
