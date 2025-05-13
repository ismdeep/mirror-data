package task

var Tasks map[string]Interface

func init() {
	Tasks = map[string]Interface{
		"alpine-linux":                  &AlpineLinux{},
		"another-redis-desktop-manager": &AnotherRedisDesktopManager{},
		"ctop":                          &Ctop{},
		"docker-compose":                &DockerCompose{},
		"electron-ssr-backup":           &ElectronSsrBackup{},
		"etcd-manager":                  &EtcdManager{},
		"git-for-windows":               &GitForWindows{},
		"go":                            &GoDev{},
		"golangci-lint":                 &GolangciLint{},
		"goose":                         &Goose{},
		"gosec":                         &GoSec{},
		"harbor":                        &Harbor{},
		"image-syncer":                  &ImageSyncer{},
		"ipfs-desktop":                  &IPFSDesktop{},
		"ipfs-kubo":                     &IPFSKubo{},
		"jetbrains":                     &JetBrains{},
		"obsidian":                      &Obsidian{},
		"openssl":                       &OpenSSL{},
		"pandoc":                        &Pandoc{},
		"python":                        &Python{},
		"rclone":                        &Rclone{},
		"ventoy":                        &Ventoy{},
		"yq":                            &Yq{},
	}
}
