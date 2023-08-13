package main

import (
	"fmt"
	"strings"

	"github.com/ismdeep/mirror-data/internal/rclone"
	"github.com/ismdeep/mirror-data/internal/store"
)

// IsCompressFile is a compress file
func IsCompressFile(path string) bool {
	return strings.Contains(path, ".tar.gz") ||
		strings.Contains(path, ".zip")
}

func main() {
	storage := store.New("adoptium", 32)
	remoteSite := "https://mirrors.tuna.tsinghua.edu.cn/Adoptium/"
	items, err := rclone.JSON("lsjson", "-R", "--http-url", remoteSite, ":http:")
	if err != nil {
		panic(err)
	}
	for _, v := range items {
		if !v.IsDir && IsCompressFile(v.Path) {
			storage.Add(v.Path, fmt.Sprintf("%v%v", remoteSite, v.Path))
		}
	}

	close(storage.C)
	storage.WG.Wait()
}
