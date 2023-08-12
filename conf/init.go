package conf

import (
	"math/rand"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	GitHubTokens []string `yaml:"ghp_list"`
}

var ROOT config

func init() {
	raw, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(raw, &ROOT); err != nil {
		panic(err)
	}

	if len(ROOT.GitHubTokens) <= 0 {
		panic("ERROR: ghp_list in config.yaml is empty")
	}
}

func RandGitHubToken() string {
	return ROOT.GitHubTokens[rand.Intn(len(ROOT.GitHubTokens))]
}
