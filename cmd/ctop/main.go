package ctop

import "github.com/ismdeep/mirror-data/internal/github"

func Run() error {
	github.FetchReleases("ctop", "bcicen", "ctop")
	return nil
}
