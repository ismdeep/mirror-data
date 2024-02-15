package task

import (
	"fmt"

	"github.com/ismdeep/mirror-data/app/data/conf"
	"github.com/ismdeep/mirror-data/app/data/global"
	"github.com/ismdeep/mirror-data/app/data/internal/rclone"
	"github.com/ismdeep/mirror-data/app/data/internal/store"
	util2 "github.com/ismdeep/mirror-data/app/data/internal/util"
	"github.com/ismdeep/mirror-data/pkg/log"
	"go.uber.org/zap"
)

type Python struct {
}

func (receiver *Python) Run() {
	storage := store.New("python", conf.Config.StorageCoroutineSize)
	defer func() {
		storage.CloseAndWait()
	}()

	remoteSite := "https://www.python.org/ftp/python/"
	versions, err := rclone.JSON("lsjson", "--http-url", remoteSite, ":http:")
	if err != nil {
		global.Errors <- err
		return
	}

	versionChan := make(chan rclone.JSONObj, 1024)
	go func() {
		defer func() {
			close(versionChan)
		}()

		for _, version := range versions {
			versionChan <- version
		}
	}()

	err = util2.CoroutineRun(conf.Config.CoroutineSize, func() error {
		for version := range versionChan {
			if !version.IsDir || !('0' <= version.Path[0] && version.Path[0] <= '9') {
				continue
			}
			items, err := rclone.JSON("lsjson", "--http-url", fmt.Sprintf("%v%v/", remoteSite, version.Path), ":http:")
			if err != nil {
				log.WithName("python").Error("failed on lsjson", zap.Error(err))
				return err
			}
			for _, v := range items {
				if !receiver.IsCompressFile(v.Path) {
					continue
				}
				link := fmt.Sprintf("%v/%v", version.Path, v.Path)
				originLink := fmt.Sprintf("%v%v/%v", remoteSite, version.Path, v.Path)
				storage.Add(link, originLink)
			}
		}

		return nil
	})
	if err != nil {
		global.Errors <- err
		return
	}
}

// IsCompressFile is compress file
func (receiver *Python) IsCompressFile(path string) bool {
	return util2.StringEndWith(path, ".tar.xz")
}
