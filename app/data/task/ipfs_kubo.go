package task

import (
	"strings"

	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type IPFSKubo struct {
}

func (receiver *IPFSKubo) Run() {
	github.FetchReleases("ipfs-kubo", "ipfs", "kubo", func(s string) bool {
		return strings.Contains(s, "-rc")
	})
}
