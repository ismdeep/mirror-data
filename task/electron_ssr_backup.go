package task

import "github.com/ismdeep/mirror-data/internal/github"

type ElectronSsrBackup struct {
}

func (receiver *ElectronSsrBackup) Run() {
	github.FetchReleases("electron-ssr-backup", "qingshuisiyuan", "electron-ssr-backup", nil)
}
