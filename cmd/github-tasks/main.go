package main

import (
	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/internal/github"
)

func main() {
	for _, task := range conf.GitHubTasks {
		github.FetchReleases(task.Bucket, task.Owner, task.Repo)
	}
}
