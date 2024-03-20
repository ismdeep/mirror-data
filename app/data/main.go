package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/ismdeep/mirror-data/app/data/internal/version"
	"github.com/ismdeep/mirror-data/app/data/task"
)

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
	}

	taskLst := os.Args[1:]
	if len(taskLst) <= 1 {
		for taskName := range task.Tasks {
			taskLst = append(taskLst, taskName)
		}
	}

	var wg sync.WaitGroup
	for _, taskName := range taskLst {
		wg.Add(1)
		go func(taskName string) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("[ERROR] failed to run task:", r)
				}
				wg.Done()
			}()
			t, ok := task.Tasks[taskName]
			if !ok {
				panic(fmt.Errorf("task not found: %v", taskName))
			}
			t.Run()

		}(taskName)
	}
	wg.Wait()
}
