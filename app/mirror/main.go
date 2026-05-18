package main

import (
	"context"
	"errors"

	"github.com/earthboundkid/versioninfo/v2"
	"github.com/ismdeep/log"
	"go.uber.org/zap"

	"github.com/ismdeep/mirror-data/app/mirror/api"
)

func main() {
	ctx := context.Background()
	log.WithContext(ctx).Info("current version info",
		zap.Any("commit_date", versioninfo.LastCommit),
		zap.String("commit_id", versioninfo.Revision))

	if err := api.Run(); err != nil {
		panic(
			errors.Join(
				errors.New("api server failed"),
				err))
	}
}
