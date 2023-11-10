package conf

import (
	"context"
	"math/rand"
	"os"

	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v3"
)

type config struct {
	CoroutineSize        int `env:"COROUTINE_SIZE,default=4"`
	StorageCoroutineSize int `env:"STORAGE_COROUTINE_SIZE,default=16"`
}

type secrets struct {
	GitHubTokens []string `yaml:"ghp_list"`
}

var Config config

// Secrets instance
var Secrets secrets

func init() {
	if err := envconfig.Process(context.TODO(), &Config); err != nil {
		panic(err)
	}

	loadYAML("secrets.yaml", &Secrets)

	if len(Secrets.GitHubTokens) <= 0 {
		panic("ghp_list in config.yaml is empty")
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
