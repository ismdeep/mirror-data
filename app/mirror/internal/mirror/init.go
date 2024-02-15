package mirror

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/ismdeep/mirror-data/app/mirror/conf"
)

// File struct
type File struct {
	Path          string    `json:"path"`
	OriginURL     string    `json:"origin_url"`
	ContentLength int64     `json:"content_length"`
	ContentType   string    `json:"content_type"`
	LastModified  time.Time `json:"last_modified"`
}

// M Mirror instance
var M *Mirror

// load bucket data
func loadBucket(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		panic(
			errors.Join(
				errors.New("failed to open bucket file"),
				err))
	}
	r := bufio.NewReaderSize(f, 1<<24) // buff size: 16MB
	for {
		text, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(io.EOF, err) {
				break
			}
			panic(
				errors.Join(
					errors.New("failed to read line from data.mirror"),
					err))
		}
		bucketName, file, err := parseText(string(text))
		if err != nil {
			panic(
				errors.Join(
					errors.New("failed to parse text"),
					err))
		}

		M.AddFile(bucketName, file)
	}

	return nil
}

func init() {
	M = NewMirror()

	if err := filepath.WalkDir(conf.ROOT.System.Data, func(path string, d fs.DirEntry, err error) error {
		// skip if is directory
		if d.IsDir() {
			return nil
		}

		// skip if file extension is not .txt
		if len(d.Name()) <= 4 || d.Name()[len(d.Name())-4:len(d.Name())] != ".txt" {
			return nil
		}

		// read and load to M
		if err := loadBucket(path); err != nil {
			return err
		}

		return nil
	}); err != nil {
		panic(
			errors.Join(
				errors.New("failed to walk dir"),
				err))
	}

	// get bucket name list
	var buckets []string
	for bucketName := range M.data {
		buckets = append(buckets, bucketName)
	}
	sort.Strings(buckets)
	M.bucketNameList = buckets
}
