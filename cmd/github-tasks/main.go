package main

import (
	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/internal/github"
)

func main() {
	for _, task := range conf.ROOT.GitHub.Tasks {
		github.FetchReleases(task.Bucket, task.Owner, task.Repo)
	}
}
