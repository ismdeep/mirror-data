package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/internal/store"
)

// FetchReleases fetch releases
func FetchReleases(bucketName string, owner string, repo string) {
	storage := store.New(bucketName, conf.Config.StorageCoroutineSize)
	cli := github.NewTokenClient(context.TODO(), conf.RandGitHubToken())
	page := 1
	for {
		releases, _, err := cli.Repositories.ListReleases(context.TODO(), owner, repo, &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			panic(err)
		}

		for _, release := range releases {
			for _, asset := range release.Assets {
				link := fmt.Sprintf("%v/%v", *release.TagName, *asset.Name)
				originLink := *asset.BrowserDownloadURL
				storage.Add(link, originLink)
			}
		}

		if len(releases) < 100 {
			break
		}
		page++
	}
	close(storage.C)
	storage.WG.Wait()
}
