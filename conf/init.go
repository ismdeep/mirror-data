package conf

import (
	"math/rand"
	"os"

	"gopkg.in/yaml.v3"
)

type secrets struct {
	GitHubTokens []string `yaml:"ghp_list"`
}

type githubTask struct {
	Bucket string `yaml:"bucket"`
	Owner  string `yaml:"owner"`
	Repo   string `yaml:"repo"`
}

// Secrets instance
var Secrets secrets

// GitHubTasks instance
var GitHubTasks []githubTask

func init() {
	loadYAML("secrets.yaml", &Secrets)
	loadYAML("github-tasks.yaml", &GitHubTasks)

	if len(Secrets.GitHubTokens) <= 0 {
		panic("ERROR: ghp_list in config.yaml is empty")
	}
}

func loadYAML(filePath string, v any) {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(raw, v); err != nil {
		panic(err)
	}
}

// RandGitHubToken get a github token by random
func RandGitHubToken() string {
	return Secrets.GitHubTokens[rand.Intn(len(Secrets.GitHubTokens))]
}
