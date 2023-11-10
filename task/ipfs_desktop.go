package task

import "github.com/ismdeep/mirror-data/internal/github"

type IPFSDesktop struct {
}

func (receiver *IPFSDesktop) Run() {
	github.FetchReleases("ipfs-desktop", "ipfs", "ipfs-desktop")
}
