package main

import "github.com/ismdeep/mirror-data/internal/github"

func main() {
	github.FetchReleases("electron-ssr-backup", "qingshuisiyuan", "electron-ssr-backup")
}
