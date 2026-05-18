package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ismdeep/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/ismdeep/mirror-data/app/data/internal/version"
	"github.com/ismdeep/mirror-data/app/data/task"
	"github.com/ismdeep/mirror-data/meta"
)

func main() {
	var metaPath string
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

			fmt.Println("All Pages:", allPages)

			raw, err := os.ReadFile(metaPath)
			if err != nil {
				return err
			}
			metaData := meta.Load(raw)

			tasks := make(map[string]task.Interface)
			if metaData.AlpineLinux.Enabled {
				tasks["alpine-linux"] = &task.AlpineLinux{}
			}
			if metaData.Go.Enabled {
				tasks["go"] = &task.GoDev{}
			}
			if metaData.JetBrains.Enabled {
				tasks["jetbrains"] = &task.JetBrains{}
			}
			if metaData.OpenSSL.Enabled {
				tasks["openssl"] = &task.OpenSSL{}
			}
			if metaData.Python.Enabled {
				tasks["python"] = &task.Python{}
			}
			for bucket, githubBucket := range metaData.GitHubBuckets {
				tasks[bucket] = task.NewGithubBucket(bucket, githubBucket.Owner, githubBucket.Repo, githubBucket.Ignored, allPages)
			}
			for _, t := range tasks {
				log.WithContext(ctx).Info("start to run task", zap.String("bucket", t.GetBucketName()))
				t.Run()
			}

			return nil
		},
	}

	m.PersistentFlags().StringVar(&metaPath, "meta", "", "path to mirror meta file")
	m.PersistentFlags().BoolVar(&allPages, "all-pages", false, "check all pages for every repository")

	_ = m.MarkPersistentFlagRequired("meta")

	if err := m.Execute(); err != nil {
		panic(err)
	}
}
