package secrets

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Secrets struct {
	GitHubTokens []string `yaml:"ghp_list"`
}

func Load(raw []byte) (*Secrets, error) {
	var secrets Secrets
	if err := yaml.Unmarshal(raw, &secrets); err != nil {
		return nil, err
	}
	return &secrets, nil
}

func LoadFromFile(filePath string) (*Secrets, error) {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return Load(raw)
}
