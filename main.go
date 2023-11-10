package main

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/ismdeep/mirror-data/cmd/adoptium"
	"github.com/ismdeep/mirror-data/cmd/alpinelinux"
	"github.com/ismdeep/mirror-data/cmd/ardm"
	"github.com/ismdeep/mirror-data/cmd/ctop"
	"github.com/ismdeep/mirror-data/cmd/dockercompose"
	"github.com/ismdeep/mirror-data/cmd/godev"
	"github.com/ismdeep/mirror-data/cmd/jetbrains"
	"github.com/ismdeep/mirror-data/cmd/nodejs"
	"github.com/ismdeep/mirror-data/cmd/openssl"
	"github.com/ismdeep/mirror-data/cmd/python"
	"github.com/ismdeep/mirror-data/pkg/log"
)

var tasks map[string]func() error

func init() {
	tasks = map[string]func() error{
		"adoptium":                      adoptium.Run,
		"alpine":                        alpinelinux.Run,
		"another-redis-desktop-manager": ardm.Run,
		"ctop":                          ctop.Run,
		"docker-compose":                dockercompose.Run,
		"go":                            godev.Run,
		"jetbrains":                     jetbrains.Run,
		"nodejs":                        nodejs.Run,
		"openssl":                       openssl.Run,
		"python":                        python.Run,
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
	errC := make(chan error, 65535)
	for _, arg := range convert(os.Args[1:]) {
		f, ok := tasks[arg]
		if !ok {
			panic(fmt.Errorf("invalid task name. [%v]", arg))
		}
		wg.Add(1)
		go func(name string, f func() error) {
			defer func() {
				log.WithName(name).Info("finished")
				wg.Done()
			}()
			if err := f(); err != nil {
				errC <- err
			}
		}(arg, f)
	}

	wg.Wait()
	close(errC)

	var errLst []error
	for err := range errC {
		errLst = append(errLst, err)
	}
	if err := errors.Join(errLst...); err != nil {
		panic(err)
	}
}
