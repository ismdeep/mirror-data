package python

import (
	"fmt"

	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/internal/rclone"
	"github.com/ismdeep/mirror-data/internal/store"
	"github.com/ismdeep/mirror-data/internal/util"
)

// IsCompressFile is compress file
func IsCompressFile(path string) bool {
	return util.StringEndWith(path, ".tar.xz")
}

func Run() error {
	storage := store.New("python", conf.Config.StorageCoroutineSize)

	remoteSite := "https://www.python.org/ftp/python/"
	versions, err := rclone.JSON("lsjson", "--http-url", remoteSite, ":http:")
	if err != nil {
		return err
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

	err = util.CoroutineRun(conf.Config.CoroutineSize, func() error {
		for version := range versionChan {
			if !version.IsDir || !('0' <= version.Path[0] && version.Path[0] <= '9') {
				continue
			}
			items, err := rclone.JSON("lsjson", "--http-url", fmt.Sprintf("%v%v/", remoteSite, version.Path), ":http:")
			if err != nil {
				fmt.Println("ERROR:", err.Error())
				return err
			}
			for _, v := range items {
				if !IsCompressFile(v.Path) {
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
		return err
	}

	storage.CloseAndWait()

	return nil
}
