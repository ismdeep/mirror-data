package main

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/ismdeep/mirror-data/app/mirror/api"
	"github.com/ismdeep/mirror-data/app/mirror/update"
	"github.com/ismdeep/mirror-data/internal/version"
	"github.com/ismdeep/mirror-data/pkg/log"
)

func main() {
	log.WithContext(context.Background()).Info("current version info", zap.String("commit_date", version.CommitDate), zap.String("commit_id", version.CommitID))
	update.StartCheckUpdateCron()

	if err := api.Run(); err != nil {
		panic(
			errors.Join(
				errors.New("api server failed"),
				err))
	}
}
