package main

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/ismdeep/mirror-data/global"
	"github.com/ismdeep/mirror-data/pkg/log"
	"github.com/ismdeep/mirror-data/task"
)

var tasks map[string]task.Interface

func init() {
	tasks = map[string]task.Interface{
		"adoptium":                      &task.Adoptium{},
		"alpine":                        &task.AlpineLinux{},
		"another-redis-desktop-manager": &task.AnotherRedisDesktopManager{},
		"ctop":                          &task.Ctop{},
		"docker-compose":                &task.DockerCompose{},
		"electron-ssr-backup":           &task.ElectronSsrBackup{},
		"etcd-manager":                  &task.EtcdManager{},
		"git-for-windows":               &task.GitForWindows{},
		"go":                            &task.GoDev{},
		"harbor":                        &task.Harbor{},
		"image-syncer":                  &task.ImageSyncer{},
		"ipfs-desktop":                  &task.IPFSDesktop{},
		"jetbrains":                     &task.JetBrains{},
		"nodejs":                        &task.NodeJS{},
		"obsidian":                      &task.Obsidian{},
		"openssl":                       &task.OpenSSL{},
		"python":                        &task.Python{},
		"rclone":                        &task.Rclone{},
		"ventoy":                        &task.Ventoy{},
	}
}

func convert(args []string) []string {
	if len(args) <= 0 || args[0] == "all" {
		var lst []string
		for name := range tasks {
			lst = append(lst, name)
		}
		return lst
	}
	return args
}

func main() {
	var wg sync.WaitGroup
	for _, arg := range convert(os.Args[1:]) {
		f, ok := tasks[arg]
		if !ok {
			panic(fmt.Errorf("invalid task name. [%v]", arg))
		}
		wg.Add(1)
		go func(name string, task task.Interface) {
			defer func() {
				log.WithName(name).Info("finished")
				wg.Done()
			}()
			task.Run()
		}(arg, f)
	}
	wg.Wait()
	close(global.Errors)

	var errLst []error
	for err := range global.Errors {
		errLst = append(errLst, err)
	}
	if len(errLst) > 0 {
		fmt.Println("================ ERROR ================")
		if err := errors.Join(errLst...); err != nil {
			fmt.Println("ERR:", err.Error())
		}
	}
}
