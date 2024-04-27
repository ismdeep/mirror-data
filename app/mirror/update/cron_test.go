package update

import (
	"context"
	"testing"

	"github.com/google/go-github/v53/github"
)

func Test_checkUpdateInfo(t *testing.T) {
	checkUpdateInfoCore(context.Background(), func(branchInfo github.Branch) error {
		t.Logf("sha1: %v", branchInfo.Commit.GetSHA())
		return nil
	})
}
