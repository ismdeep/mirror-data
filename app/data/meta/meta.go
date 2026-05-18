package meta

import (
	"os"

	"gopkg.in/yaml.v3"
)

type GitHubBucket struct {
	Owner   string `yaml:"owner"`
	Repo    string `yaml:"repo"`
	Ignored string `yaml:"ignored"`
}

type AlpineLinux struct {
	Enabled bool `yaml:"enabled"`
}

type Go struct {
	Enabled bool `yaml:"enabled"`
}

type JetBrains struct {
	Enabled bool `yaml:"enabled"`
}

type OpenSSL struct {
	Enabled bool `yaml:"enabled"`
}

type Python struct {
	Enabled bool `yaml:"enabled"`
}

type Meta struct {
	GitHubBuckets map[string]GitHubBucket `yaml:"github_buckets"`
	AlpineLinux   AlpineLinux             `yaml:"alpine_linux"`
	Go            Go                      `yaml:"go"`
	JetBrains     JetBrains               `yaml:"jetbrains"`
	OpenSSL       OpenSSL                 `yaml:"openssl"`
	Python        Python                  `yaml:"python"`
}

func Load(data []byte) (*Meta, error) {
	var meta Meta
	if err := yaml.Unmarshal(data, &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

func LoadFromFile(filePath string) (*Meta, error) {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return Load(raw)
}
