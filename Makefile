all: ctop docker-compose harbor image-syncer git-for-windows electron-ssr-backup jetbrains

ctop:
	go run ./cmd/ctop

docker-compose:
	go run ./cmd/docker-compose

harbor:
	go run ./cmd/harbor

image-syncer:
	go run ./cmd/image-syncer

git-for-windows:
	go run ./cmd/git-for-windows

electron-ssr-backup:
	go run ./cmd/electron-ssr-backup

jetbrains:
	go run ./cmd/jetbrains
