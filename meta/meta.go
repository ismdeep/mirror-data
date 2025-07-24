package meta

import "gopkg.in/yaml.v3"

type GitHubBucket struct {
	Owner   string `yaml:"owner"`
	Repo    string `yaml:"repo"`
	Ignored string `yaml:"ignored"`
}

type Meta struct {
	GitHubBuckets map[string]GitHubBucket `yaml:"github_buckets"`
}

func Load(data []byte) Meta {
	var meta Meta
	if err := yaml.Unmarshal(data, &meta); err != nil {
		panic(err)
	}
	return meta
}
