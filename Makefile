all: ctop docker-compose harbor image-syncer git-for-windows electron-ssr-backup jetbrains ipfs-desktop another-redis-desktop-manager etcd-manager

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

ipfs-desktop:
	go run ./cmd/ipfs-desktop

another-redis-desktop-manager:
	go run ./cmd/another-redis-desktop-manager

etcd-manager:
	go run ./cmd/etcd-manager
