package conf

import (
	"math/rand"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	GitHub struct {
		Tokens []string `yaml:"tokens"`
		Tasks  []struct {
			Bucket string `yaml:"bucket"`
			Owner  string `yaml:"owner"`
			Repo   string `yaml:"repo"`
		} `yaml:"tasks"`
	} `yaml:"github"`
}

// ROOT instance
var ROOT config

func init() {
	raw, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(raw, &ROOT); err != nil {
		panic(err)
	}

	if len(ROOT.GitHub.Tokens) <= 0 {
		panic("ERROR: ghp_list in config.yaml is empty")
	}
}

// RandGitHubToken get a github token by random
func RandGitHubToken() string {
	return ROOT.GitHub.Tokens[rand.Intn(len(ROOT.GitHub.Tokens))]
}
