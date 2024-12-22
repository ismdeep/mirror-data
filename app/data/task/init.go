package task

var Tasks map[string]Interface

func init() {
	Tasks = map[string]Interface{
		"electron-ssr-backup":           &ElectronSsrBackup{},
		"etcd-manager":                  &EtcdManager{},
		"git-for-windows":               &GitForWindows{},
		"go":                            &GoDev{},
		"goose":                         &Goose{},
		"gosec":                         &GoSec{},
		"harbor":                        &Harbor{},
		"ipfs-desktop":                  &IPFSDesktop{},
		"jetbrains":                     &JetBrains{},
		"obsidian":                      &Obsidian{},
		"openssl":                       &OpenSSL{},
		"rclone":                        &Rclone{},
		"ventoy":                        &Ventoy{},
		"ctop":                          &Ctop{},
		"another-redis-desktop-manager": &AnotherRedisDesktopManager{},
		"docker-compose":                &DockerCompose{},
		"image-syncer":                  &ImageSyncer{},
		"ipfs-kubo":                     &IPFSKubo{},
		"pandoc":                        &Pandoc{},
		"yq":                            &Yq{},
	}
}
