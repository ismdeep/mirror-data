package update

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v53/github"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/ismdeep/mirror-data/app/mirror/conf"
	"github.com/ismdeep/mirror-data/internal/version"
	"github.com/ismdeep/mirror-data/pkg/log"
	"github.com/ismdeep/mirror-data/pkg/notification"
)

func checkUpdateInfoCore(ctx context.Context, callback func(branchInfo github.Branch) error) {
	cli := github.NewClient(&http.Client{})
	branchInfo, _, err := cli.Repositories.GetBranch(ctx, "ismdeep", "mirror-data", "main", true)
	if err != nil {
		log.WithContext(ctx).Error("failed to get branch info", zap.Error(err))
		return
	}

	if branchInfo.Commit.GetSHA() != version.CommitID {
		if err := callback(*branchInfo); err != nil {
			log.WithContext(ctx).Error("callback failed", zap.Error(err))
		}
	}
}

func CheckUpdateInfo(ctx context.Context) {
	defer func() {
		log.WithContext(ctx).Info("check update info finished")
	}()
	log.WithContext(ctx).Info("start check update info")
	checkUpdateInfoCore(ctx, func(branchInfo github.Branch) error {
		msg := fmt.Sprintf("[UPDATE] New version detected. Please update to the latest version in a timely manner. %v", conf.ROOT.Server.Site)
		if err := notification.RelayMsg([]string{conf.ROOT.Notification.Relay.Endpoint}, conf.ROOT.Notification.Relay.Auth, msg); err != nil {
			log.WithContext(ctx).Error("relay msg failed", zap.Error(err))
			return err
		}
		return nil
	})
}

func StartCheckUpdateCron() {
	CheckUpdateInfo(log.NewTraceContext(uuid.NewString()))
	go func() {
		ticker := time.NewTicker(30 * time.Minute)
		for range ticker.C {
			CheckUpdateInfo(log.NewTraceContext(uuid.NewString()))
		}
	}()
}
