package task

import "github.com/ismdeep/mirror-data/internal/github"

type Obsidian struct {
}

func (receiver *Obsidian) Run() {
	github.FetchReleases("obsidian", "obsidianmd", "obsidian-releases")
}
