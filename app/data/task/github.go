package task

import (
	"regexp"

	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type GithubBucket struct {
	bucket                     string
	owner                      string
	repo                       string
	ignored                    string
	checkAllGitHubReleasePages bool
}

func NewGithubBucket(bucket string, owner string, repo string, ignored string, checkAllGitHubReleasePages bool) *GithubBucket {
	return &GithubBucket{
		bucket:                     bucket,
		owner:                      owner,
		repo:                       repo,
		ignored:                    ignored,
		checkAllGitHubReleasePages: checkAllGitHubReleasePages,
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

	github.FetchReleases(receiver.bucket, receiver.owner, receiver.repo, f, receiver.checkAllGitHubReleasePages)
}

func (receiver *GithubBucket) GetBucketName() string {
	return receiver.bucket
}
