package conf

import (
	"context"
	"math/rand"

	"github.com/sethvargo/go-envconfig"

	"github.com/ismdeep/mirror-data/app/data/secrets"
)

type config struct {
	CoroutineSize        int `env:"COROUTINE_SIZE,default=4"`
	StorageCoroutineSize int `env:"STORAGE_COROUTINE_SIZE,default=16"`
}

var Config config

// Secrets instance
var Secrets *secrets.Secrets

func Init() error {
	if err := envconfig.Process(context.TODO(), &Config); err != nil {
		return err
	}
	return nil
}

// RandGitHubToken get a github token by random
func RandGitHubToken() string {
	return Secrets.GitHubTokens[rand.Intn(len(Secrets.GitHubTokens))]
}
