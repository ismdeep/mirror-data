package main

import (
	"context"
	"errors"

	"github.com/earthboundkid/versioninfo/v2"
	"go.uber.org/zap"

	"github.com/ismdeep/mirror-data/app/mirror/api"
	"github.com/ismdeep/mirror-data/app/mirror/update"
	"github.com/ismdeep/mirror-data/pkg/log"
)

func main() {
	log.WithContext(context.Background()).Info("current version info",
		zap.Any("commit_date", versioninfo.LastCommit),
		zap.String("commit_id", versioninfo.Revision))
	update.StartCheckUpdateCron()

	if err := api.Run(); err != nil {
		panic(
			errors.Join(
				errors.New("api server failed"),
				err))
	}
}
