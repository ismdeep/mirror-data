package task

import (
	"fmt"
	"strings"

	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/global"
	"github.com/ismdeep/mirror-data/internal/rclone"
	"github.com/ismdeep/mirror-data/internal/store"
	"github.com/ismdeep/mirror-data/internal/util"
	"github.com/ismdeep/mirror-data/pkg/log"
	"go.uber.org/zap"
)

type NodeJS struct {
}

func (receiver *NodeJS) Run() {
	storage := store.New("nodejs", conf.Config.StorageCoroutineSize)
	defer func() {
		storage.CloseAndWait()
	}()

	versions, err := rclone.JSON("lsjson", "--http-url", "https://nodejs.org/dist/", ":http:")
	if err != nil {
		global.Errors <- err
		return
	}

	versionChan := make(chan rclone.JSONObj, 1024)
	go func() {
		for _, version := range versions {
			versionChan <- version
		}
		close(versionChan)
	}()

	err = util.CoroutineRun(conf.Config.CoroutineSize, func() error {
		for version := range versionChan {
			if !version.IsDir || receiver.isIgnoredPath(version.Path) {
				continue
			}
			items, err := rclone.JSON("lsjson", "--http-url", fmt.Sprintf("https://nodejs.org/dist/%v/", version.Path), ":http:")
			if err != nil {
				log.WithName("nodejs").Error("failed on lsjson", zap.Error(err))
				return err
			}
			for _, v := range items {
				if !receiver.isCompressFile(v.Path) {
					continue
				}
				link := fmt.Sprintf("%v/%v", version.Path, v.Path)
				originLink := fmt.Sprintf("https://nodejs.org/dist/%v/%v", version.Path, v.Path)
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

// isCompressFile 是否是压缩包
func (receiver *NodeJS) isCompressFile(path string) bool {
	return strings.Contains(path, ".tar.gz") ||
		strings.Contains(path, ".tar.xz") ||
		strings.Contains(path, ".zip") ||
		strings.Contains(path, ".7z")
}

// isIgnoredPath 是否被忽略
func (receiver *NodeJS) isIgnoredPath(path string) bool {
	return strings.Contains(path, "latest") ||
		strings.Contains(path, "npm") ||
		strings.Contains(path, "patch") ||
		strings.Contains(path, "-isaacs-manual")
}
