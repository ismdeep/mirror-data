package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ismdeep/mirror-data/global"
	"github.com/ismdeep/mirror-data/internal/version"
	"github.com/ismdeep/mirror-data/pkg/log"
	"github.com/ismdeep/mirror-data/task"
)

var tasks map[string]task.Interface

func init() {
	tasks = map[string]task.Interface{
		"another-redis-desktop-manager": &task.AnotherRedisDesktopManager{},
		"ctop":                          &task.Ctop{},
		"docker-compose":                &task.DockerCompose{},
		"etcd-manager":                  &task.EtcdManager{},
		"git-for-windows":               &task.GitForWindows{},
		"go":                            &task.GoDev{},
		"harbor":                        &task.Harbor{},
		"image-syncer":                  &task.ImageSyncer{},
		"ipfs-desktop":                  &task.IPFSDesktop{},
		"ipfs-kubo":                     &task.IPFSKubo{},
		"jetbrains":                     &task.JetBrains{},
		"obsidian":                      &task.Obsidian{},
		"openssl":                       &task.OpenSSL{},
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

	if version.Version != "" {
		fmt.Println("MIRROR DATA")
		fmt.Println("Version:", version.Version)
		fmt.Println("OS:", version.OS)
		fmt.Println("Arch:", version.Arch)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("Commit ID:", version.CommitID)
		fmt.Println("Commit Date:", version.CommitDate)
		fmt.Println()
		time.Sleep(300 * time.Millisecond)
	}

	var wg sync.WaitGroup
	lst := convert(os.Args[1:])
	for _, arg := range lst {
		t, ok := tasks[arg]
		if !ok {
			panic(fmt.Errorf("invalid task name. [%v]", arg))
		}
		wg.Add(1)
		go func(name string, t task.Interface) {
			t.Run()
			log.WithName(name).Info("finished")
			wg.Done()
		}(arg, t)
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
