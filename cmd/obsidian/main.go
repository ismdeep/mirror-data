package main

import "github.com/ismdeep/mirror-data/internal/github"

func main() {
	github.FetchReleases("obsidian", "obsidianmd", "obsidian-releases")
}
