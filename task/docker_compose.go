package task

import (
	"strings"

	"github.com/ismdeep/mirror-data/internal/github"
)

type DockerCompose struct {
}

func (receiver *DockerCompose) Run() {
	github.FetchReleases("docker-compose", "docker", "compose", func(s string) bool {
		return strings.Contains(s, "-rc")
	})
}
