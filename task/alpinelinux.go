package task

import (
	"fmt"

	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/global"
	"github.com/ismdeep/mirror-data/internal/rclone"
	"github.com/ismdeep/mirror-data/internal/store"
	"github.com/ismdeep/mirror-data/internal/util"
	"github.com/ismdeep/mirror-data/pkg/log"
	"go.uber.org/zap"
)

type AlpineLinux struct {
}

func (receiver *AlpineLinux) Run() {
	storage := store.New("alpine-linux", conf.Config.StorageCoroutineSize)

	remoteSite := "https://dl-cdn.alpinelinux.org/alpine"

	versionsResp, err := rclone.JSON("lsjson", "--http-url", remoteSite, ":http:")
	if err != nil {
		global.Errors <- err
	}

	versionChan := make(chan string, 1024)
	go func() {
		defer func() {
			close(versionChan)
		}()

		for _, vResp := range versionsResp {
			if util.IsAlpineVersion(vResp.Path) {
				versionChan <- vResp.Path
			}
		}
	}()

	err = util.CoroutineRun(conf.Config.CoroutineSize, func() error {
		for version := range versionChan {
			items, err := rclone.JSON("lsjson", "-R", "--http-url", fmt.Sprintf("%v/%v/releases/", remoteSite, version), ":http:")
			if err != nil {
				log.WithName("alpine-linux").Error("failed on lsjson", zap.Error(err))
				return err
			}
			for _, v := range items {
				if v.IsDir {
					continue
				}

				if util.StringEndWith(v.Path, ".iso") {
					storage.Add(
						fmt.Sprintf("%v/%v", version, v.Path),
						fmt.Sprintf("%v/%v/releases/%v", remoteSite, version, v.Path))
				}
			}
		}
		return nil
	})
	if err != nil {
		global.Errors <- err
		return
	}

	storage.CloseAndWait()
}
