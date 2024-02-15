package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"github.com/ismdeep/mirror-data/app/data/conf"
	"github.com/ismdeep/mirror-data/app/data/global"
	"github.com/ismdeep/mirror-data/app/data/internal/store"
	"github.com/ismdeep/mirror-data/pkg/log"
	"go.uber.org/zap"
)

// FetchReleases fetch releases
func FetchReleases(bucketName string, owner string, repo string, ignoredFunc func(s string) bool) {
	storage := store.New(bucketName, conf.Config.StorageCoroutineSize)
	defer func() {
		storage.CloseAndWait()
	}()

	token := conf.RandGitHubToken()
	cli := github.NewTokenClient(context.TODO(), token)
	page := 1
	for {
		log.WithName(bucketName).Info(bucketName, zap.Any("page", page))
		releases, _, err := cli.Repositories.ListReleases(context.TODO(), owner, repo, &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			log.WithContext(context.TODO()).Error("failed to list releases",
				zap.String("token", token),
				zap.Error(err))
			global.Errors <- err
			break
		}

		for _, release := range releases {
			for _, asset := range release.Assets {
				link := fmt.Sprintf("%v/%v", *release.TagName, *asset.Name)
				if ignoredFunc != nil && ignoredFunc(link) {
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
