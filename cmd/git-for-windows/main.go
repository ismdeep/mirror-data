package main

import "github.com/ismdeep/mirror-data/internal/github"

func main() {
	github.FetchReleases("git-for-windows", "git-for-windows", "git")
}
