package task

import (
	"fmt"
	"strings"

	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/global"
	"github.com/ismdeep/mirror-data/internal/rclone"
	"github.com/ismdeep/mirror-data/internal/store"
)

type Adoptium struct {
}

// isCompressFile is a compress file
func (receiver *Adoptium) isCompressFile(path string) bool {
	return strings.Contains(path, ".tar.gz") ||
		strings.Contains(path, ".zip")
}

func (receiver *Adoptium) Run() {
	storage := store.New("adoptium", conf.Config.StorageCoroutineSize)
	defer func() {
		storage.CloseAndWait()
	}()

	remoteSite := "https://mirrors.tuna.tsinghua.edu.cn/Adoptium/"
	items, err := rclone.JSON("lsjson", "-R", "--http-url", remoteSite, ":http:")
	if err != nil {
		global.Errors <- err
		return
	}
	for _, v := range items {
		if !v.IsDir && receiver.isCompressFile(v.Path) {
			storage.Add(v.Path, fmt.Sprintf("%v%v", remoteSite, v.Path))
		}
	}
}
