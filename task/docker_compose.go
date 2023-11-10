package task

import "github.com/ismdeep/mirror-data/internal/github"

type DockerCompose struct {
}

func (receiver *DockerCompose) Run() {
	github.FetchReleases("docker-compose", "docker", "compose")
}
