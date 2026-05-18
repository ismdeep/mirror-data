package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"github.com/ismdeep/log"
	"go.uber.org/zap"

	"github.com/ismdeep/mirror-data/app/data/conf"
	"github.com/ismdeep/mirror-data/app/data/global"
	"github.com/ismdeep/mirror-data/app/data/internal/store"
)

// FetchReleases fetch releases
func FetchReleases(bucketName string, owner string, repo string, ignoredFunc func(s string) bool, checkAllPages bool) {
	ctx := context.Background()

	storage := store.New(bucketName, conf.Config.StorageCoroutineSize)
	defer func() {
		storage.CloseAndWait()
	}()

	token := conf.RandGitHubToken()
	cli := github.NewTokenClient(context.TODO(), token)
	page := 1
	for {
		log.WithContext(ctx).Info("fetching",
			zap.Any("page", page),
			zap.String("bucket", bucketName),
			zap.String("owner", owner),
			zap.String("repo", repo))
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
					log.WithContext(ctx).Debug("ignored", zap.String("bucket", bucketName),
						zap.String("link", link))
					continue
				}
				originLink := *asset.BrowserDownloadURL
				storage.Add(link, originLink)
			}
		}

		if !checkAllPages || len(releases) < 100 {
			break
		}
		page++
	}
}
