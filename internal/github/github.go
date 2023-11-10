package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/global"
	"github.com/ismdeep/mirror-data/internal/store"
	"github.com/ismdeep/mirror-data/pkg/log"
	"go.uber.org/zap"
)

// FetchReleases fetch releases
func FetchReleases(bucketName string, owner string, repo string, ignoredFunc func(s string) bool) {
	storage := store.New(bucketName, conf.Config.StorageCoroutineSize)
	defer func() {
		storage.CloseAndWait()
	}()

	cli := github.NewTokenClient(context.TODO(), conf.RandGitHubToken())
	page := 1
	for {
		releases, _, err := cli.Repositories.ListReleases(context.TODO(), owner, repo, &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			global.Errors <- err
			continue
		}

		for _, release := range releases {
			for _, asset := range release.Assets {
				link := fmt.Sprintf("%v/%v", *release.TagName, *asset.Name)
				if ignoredFunc(link) {
					log.WithName(bucketName).Debug("ignored", zap.String("link", link))
					continue
				}
				originLink := *asset.BrowserDownloadURL
				storage.Add(link, originLink)
			}
		}

		if len(releases) < 100 {
			break
		}
		page++
	}
}
