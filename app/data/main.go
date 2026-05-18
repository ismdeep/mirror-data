package main

import (
	"context"
	"fmt"

	"github.com/ismdeep/log"
	"github.com/kopeisec/fp"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/ismdeep/mirror-data/app/data/conf"
	"github.com/ismdeep/mirror-data/app/data/internal/version"
	"github.com/ismdeep/mirror-data/app/data/meta"
	"github.com/ismdeep/mirror-data/app/data/secrets"
	"github.com/ismdeep/mirror-data/app/data/task"
)

type TaskItem struct {
	Name string
	Task task.Interface
}

func main() {
	var metaPath string
	var secretsPath string
	var repos []string
	var allPages bool

	m := cobra.Command{
		Use:   "mirror",
		Short: "mirror",
		Long:  "mirror",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

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

			fmt.Println("metaPath:", metaPath)
			fmt.Println("secretsPath:", secretsPath)
			fmt.Println("allPages:", allPages)
			fmt.Println("repos:", repos)

			// load meta
			metaData, err := meta.LoadFromFile(metaPath)
			if err != nil {
				return err
			}

			// load secrets
			secretsData, err := secrets.LoadFromFile(secretsPath)
			if err != nil {
				return err
			}

			// load config
			if err := conf.Init(); err != nil {
				return err
			}
			conf.Secrets = secretsData

			// load tasks
			var tasks []TaskItem
			if metaData.AlpineLinux.Enabled {
				tasks = append(tasks, TaskItem{Name: "alpine-linux", Task: &task.AlpineLinux{}})
			}
			if metaData.Go.Enabled {
				tasks = append(tasks, TaskItem{Name: "go", Task: &task.GoDev{}})
			}
			if metaData.JetBrains.Enabled {
				tasks = append(tasks, TaskItem{Name: "jetbrains", Task: &task.JetBrains{}})
			}
			if metaData.OpenSSL.Enabled {
				tasks = append(tasks, TaskItem{Name: "openssl", Task: &task.OpenSSL{}})
			}
			if metaData.Python.Enabled {
				tasks = append(tasks, TaskItem{Name: "python", Task: &task.Python{}})
			}
			for bucket, githubBucket := range metaData.GitHubBuckets {
				tasks = append(tasks, TaskItem{
					Name: bucket,
					Task: task.NewGithubBucket(bucket, githubBucket.Owner, githubBucket.Repo, githubBucket.Ignored, allPages),
				})
			}

			// selected tasks
			if len(repos) != 0 {
				tasks = fp.Wrap(tasks).Filter(func(item TaskItem) bool {
					return StringSliceContains(repos, item.Name)
				})
			}

			// run tasks
			for _, t := range tasks {
				log.WithContext(ctx).Info("start to run task", zap.String("bucket", t.Task.GetBucketName()))
				t.Task.Run()
			}

			return nil
		},
	}

	m.PersistentFlags().StringVar(&metaPath, "meta", "", "path to mirror meta file")
	m.PersistentFlags().StringVar(&secretsPath, "secrets", "", "path to mirror secrets file")
	m.PersistentFlags().BoolVar(&allPages, "all-pages", false, "check all pages for every repository")
	m.PersistentFlags().StringSliceVar(&repos, "repo", []string{}, "mirror repos to mirror")

	_ = m.MarkPersistentFlagRequired("meta")
	_ = m.MarkPersistentFlagRequired("secrets")

	if err := m.Execute(); err != nil {
		panic(err)
	}
}

func StringSliceContains(slice []string, s string) bool {
	return len(fp.Wrap(slice).Filter(func(t string) bool {
		return s == t
	})) > 0
}
