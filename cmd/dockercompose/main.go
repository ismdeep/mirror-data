package dockercompose

import "github.com/ismdeep/mirror-data/internal/github"

func Run() error {
	github.FetchReleases("docker-compose", "docker", "compose")
	return nil
}
