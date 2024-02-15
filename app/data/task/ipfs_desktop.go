package task

import (
	"strings"

	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type IPFSDesktop struct {
}

func (receiver *IPFSDesktop) Run() {
	github.FetchReleases("ipfs-desktop", "ipfs", "ipfs-desktop", func(s string) bool {
		return strings.Contains(s, "-rc.") || strings.Contains(s, "-beta.")
	})
}
