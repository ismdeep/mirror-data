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

			tasks := task.Tasks

			if metaPath != "" {
				raw, err := os.ReadFile(metaPath)
				if err != nil {
					return err
				}
				metaMap := meta.Load(raw)
				for bucket, githubBucket := range metaMap.GitHubBuckets {
					tasks[bucket] = task.NewGithubBucket(bucket, githubBucket.Owner, githubBucket.Repo, githubBucket.Ignored)
				}
			}

			for _, t := range tasks {
				log.WithContext(ctx).Info("start to run task", zap.String("bucket", t.GetBucketName()))
				t.Run()
			}

			return nil
		},
	}

	m.PersistentFlags().StringVar(&metaPath, "meta", "", "path to mirror meta file")

	if err := m.Execute(); err != nil {
		panic(err)
	}
}
