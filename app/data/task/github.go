package task

import (
	"regexp"

	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type GithubBucket struct {
	bucket  string
	owner   string
	repo    string
	ignored string
}

func NewGithubBucket(bucket string, owner string, repo string, ignored string) *GithubBucket {
	return &GithubBucket{
		bucket:  bucket,
		owner:   owner,
		repo:    repo,
		ignored: ignored,
	}
}

func (receiver *GithubBucket) Run() {
	var f func(string) bool
	if receiver.ignored != "" {
		f = func(s string) bool {
			re := regexp.MustCompile(receiver.ignored)
			return re.MatchString(s)
		}
	}

	github.FetchReleases(receiver.bucket, receiver.owner, receiver.repo, f)
}

func (receiver *GithubBucket) GetBucketName() string {
	return receiver.bucket
}
