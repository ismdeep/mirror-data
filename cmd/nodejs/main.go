package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ismdeep/mirror-data/internal/rclone"
	"github.com/ismdeep/mirror-data/internal/store"
)

// IsCompressFile 是否是压缩包
func IsCompressFile(path string) bool {
	return strings.Contains(path, ".tar.gz") ||
		strings.Contains(path, ".tar.xz") ||
		strings.Contains(path, ".zip") ||
		strings.Contains(path, ".7z")
}

// IsIgnoredPath 是否被忽略
func IsIgnoredPath(path string) bool {
	return strings.Contains(path, "latest") ||
		strings.Contains(path, "npm") ||
		strings.Contains(path, "patch") ||
		strings.Contains(path, "-isaacs-manual")
}

func main() {
	storage := store.New("nodejs", 32)

	versions, err := rclone.JSON("lsjson", "--http-url", "https://nodejs.org/dist/", ":http:")
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
			if !version.IsDir || IsIgnoredPath(version.Path) {
				return
			}
			items, err := rclone.JSON("lsjson", "--http-url", fmt.Sprintf("https://nodejs.org/dist/%v/", version.Path), ":http:")
			if err != nil {
				fmt.Println("ERROR:", err.Error())
				return
			}
			for _, v := range items {
				if !IsCompressFile(v.Path) {
					continue
				}
				link := fmt.Sprintf("%v/%v", version.Path, v.Path)
				originLink := fmt.Sprintf("https://nodejs.org/dist/%v/%v", version.Path, v.Path)
				storage.Add(link, originLink)
			}
		}(version)
	}

	wg.Wait()
	close(storage.C)
	storage.WG.Wait()
}
