package main

import (
	"fmt"
	"sync"

	"github.com/ismdeep/mirror-data/internal/rclone"
	"github.com/ismdeep/mirror-data/internal/store"
	"github.com/ismdeep/mirror-data/internal/util"
)

// IsCompressFile is compress file
func IsCompressFile(path string) bool {
	return util.StringEndWith(path, ".tar.xz")
}

func main() {
	storage := store.New("python", 16)

	remoteSite := "https://www.python.org/ftp/python/"
	versions, err := rclone.JSON("lsjson", "--http-url", remoteSite, ":http:")
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, version := range versions {
		wg.Add(1)
		go func(version rclone.JSONObj) {
			defer func() {
				wg.Done()
			}()
			if !version.IsDir || !('0' <= version.Path[0] && version.Path[0] <= '9') {
				return
			}
			items, err := rclone.JSON("lsjson", "--http-url", fmt.Sprintf("%v%v/", remoteSite, version.Path), ":http:")
			if err != nil {
				fmt.Println("ERROR:", err.Error())
				return
			}
			for _, v := range items {
				if !IsCompressFile(v.Path) {
					continue
				}
				link := fmt.Sprintf("%v/%v", version.Path, v.Path)
				originLink := fmt.Sprintf("%v%v/%v", remoteSite, version.Path, v.Path)
				storage.Add(link, originLink)
			}
		}(version)

	}

	wg.Wait()
	storage.CloseAndWait()
}
