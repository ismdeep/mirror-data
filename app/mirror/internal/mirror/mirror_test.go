package mirror

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMirror_AddFile(t *testing.T) {
	testM := NewMirror()
	testM.AddFile("go", File{
		Path:          "go1.20.5.linux-amd64.tar.gz",
		OriginURL:     "https://go.dev/dl/go1.20.5.linux-amd64.tar.gz",
		ContentLength: 100203442,
		ContentType:   "application/x-gzip",
		LastModified:  time.Unix(1688394082, 0),
	})

	type args struct {
		bucketName string
		file       File
	}
	tests := []struct {
		name string
		args args
		m    *Mirror
		want bool
	}{
		{
			name: "",
			args: args{
				bucketName: "go",
				file: File{
					Path:          "go1.20.5.linux-amd64.tar.gz",
					OriginURL:     "https://go.dev/dl/go1.20.5.linux-amd64.tar.gz",
					ContentLength: 100203442,
					ContentType:   "application/x-gzip",
					LastModified:  time.Unix(1688394082, 0),
				},
			},
			m:    NewMirror(),
			want: true,
		},
		{
			name: "",
			args: args{
				bucketName: "go",
				file: File{
					Path:          "go1.20.5.linux-amd64.tar.gz",
					OriginURL:     "https://go.dev/dl/go1.20.5.linux-amd64.tar.gz",
					ContentLength: 100203442,
					ContentType:   "application/x-gzip",
					LastModified:  time.Unix(1688394082, 0),
				},
			},
			m:    testM,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.AddFile(tt.args.bucketName, tt.args.file); got != tt.want {
				t.Errorf("AddFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMirror_FileExists(t *testing.T) {
	m := NewMirror()
	m.AddFile("go", File{
		Path:          "go1.20.5.linux-amd64.tar.gz",
		OriginURL:     "https://go.dev/dl/go1.20.5.linux-amd64.tar.gz",
		ContentLength: 100203442,
		ContentType:   "application/x-gzip",
		LastModified:  time.Unix(1688394082, 0),
	})

	assert.True(t, m.FileExists("go", "go1.20.5.linux-amd64.tar.gz"))
	assert.False(t, m.FileExists("go", "go1.20.5.linux-amd64.tar.gz-not-exists"))
}

func TestMirror_ListBuckets(t *testing.T) {
	m := NewMirror()
	m.AddFile("go", File{
		Path:          "go1.20.5.linux-amd64.tar.gz",
		OriginURL:     "https://go.dev/dl/go1.20.5.linux-amd64.tar.gz",
		ContentLength: 100203442,
		ContentType:   "application/x-gzip",
		LastModified:  time.Unix(1688394082, 0),
	})
	m.AddFile("go", File{
		Path:          "go1.20.6.linux-amd64.tar.gz",
		OriginURL:     "https://go.dev/dl/go1.20.6.linux-amd64.tar.gz",
		ContentLength: 100203442,
		ContentType:   "application/x-gzip",
		LastModified:  time.Unix(1688394082, 0),
	})

	m.AddFile("kubernetes", File{
		Path:          "kubelet",
		OriginURL:     "https://kubernetes.org/kubelet/v1.23.6/amd64/kubelet",
		ContentLength: 100203442,
		ContentType:   "application/x-gzip",
		LastModified:  time.Unix(1688394082, 0),
	})

	buckets := m.ListBuckets()
	for i, bucket := range buckets {
		t.Logf("%v: %v", i, bucket)
	}
}

func TestMirror_ListFiles(t *testing.T) {
	m := NewMirror()
	m.AddFile("go", File{
		Path:          "go1.20.5.linux-amd64.tar.gz",
		OriginURL:     "https://go.dev/dl/go1.20.5.linux-amd64.tar.gz",
		ContentLength: 100203442,
		ContentType:   "application/x-gzip",
		LastModified:  time.Unix(1688394082, 0),
	})
	m.AddFile("go", File{
		Path:          "go1.20.6.linux-amd64.tar.gz",
		OriginURL:     "https://go.dev/dl/go1.20.6.linux-amd64.tar.gz",
		ContentLength: 100203442,
		ContentType:   "application/x-gzip",
		LastModified:  time.Unix(1688394082, 0),
	})

	m.AddFile("kubernetes", File{
		Path:          "kubelet",
		OriginURL:     "https://kubernetes.org/kubelet/v1.23.6/amd64/kubelet",
		ContentLength: 100203442,
		ContentType:   "application/x-gzip",
		LastModified:  time.Unix(1688394082, 0),
	})

	files := m.ListFiles("go")
	for i, file := range files {
		t.Logf("%v: %v", i, file.Path)
	}
}
