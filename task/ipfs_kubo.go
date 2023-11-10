package task

import "github.com/ismdeep/mirror-data/internal/github"

type IPFSKubo struct {
}

func (receiver *IPFSKubo) Run() {
	github.FetchReleases("ipfs-kubo", "ipfs", "kubo")
}
