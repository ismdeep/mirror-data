package main

import (
	"fmt"

	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/internal/rclone"
	"github.com/ismdeep/mirror-data/internal/store"
	"github.com/ismdeep/mirror-data/internal/util"
)

func main() {
	storage := store.New("alpine-linux", conf.Config.StorageCoroutineSize)

	remoteSite := "https://dl-cdn.alpinelinux.org/alpine"
	versions := []string{
		"v3.0",
		"v3.1",
		"v3.2",
		"v3.3",
		"v3.4",
		"v3.5",
		"v3.6",
		"v3.7",
		"v3.8",
		"v3.9",
		"v3.10",
		"v3.11",
		"v3.12",
		"v3.13",
		"v3.14",
		"v3.15",
		"v3.16",
		"v3.17",
		"v3.18",
	}

	versionChan := make(chan string, 1024)
	go func() {
		for _, version := range versions {
			versionChan <- version
		}
		close(versionChan)
	}()

	err := util.CoroutineRun(conf.Config.CoroutineSize, func() error {
		for version := range versionChan {
			items, err := rclone.JSON("lsjson", "-R", "--http-url", fmt.Sprintf("%v/%v/releases/", remoteSite, version), ":http:")
			if err != nil {
				fmt.Println("ERROR:", err.Error())
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
		panic(err)
	}

	storage.CloseAndWait()
}
