package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type Protobuf struct {
}

func (receiver *Protobuf) Run() {
	github.FetchReleases("protobuf", "protocolbuffers", "protobuf", nil)
}
