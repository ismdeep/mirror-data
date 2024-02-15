package mirror

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseText(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   File
		wantErr bool
	}{
		{
			// ok
			name: "",
			args: args{
				s: "go|go1.20.5.linux-amd64.tar.gz|https://go.dev/dl/go1.20.5.linux-amd64.tar.gz|100203442|application/x-gzip|1688394082",
			},
			want: "go",
			want1: File{
				Path:          "go1.20.5.linux-amd64.tar.gz",
				OriginURL:     "https://go.dev/dl/go1.20.5.linux-amd64.tar.gz",
				ContentLength: 100203442,
				ContentType:   "application/x-gzip",
				LastModified:  time.Unix(1688394082, 0),
			},
			wantErr: false,
		},
		{
			// invalid text
			name: "go|go1.20.5.linux-amd64.tar.gz",
			args: args{
				s: "",
			},
			want: "",
			want1: File{
				Path:          "",
				OriginURL:     "",
				ContentLength: 0,
				ContentType:   "",
				LastModified:  time.Time{},
			},
			wantErr: true,
		},
		{
			// invalid content-length
			name: "",
			args: args{
				s: "go|go1.20.5.linux-amd64.tar.gz|https://go.dev/dl/go1.20.5.linux-amd64.tar.gz|invalid-content-length|application/x-gzip|1688394082",
			},
			want: "",
			want1: File{
				Path:          "",
				OriginURL:     "",
				ContentLength: 0,
				ContentType:   "",
				LastModified:  time.Time{},
			},
			wantErr: true,
		},
		{
			// invalid last-modified
			name: "",
			args: args{
				s: "go|go1.20.5.linux-amd64.tar.gz|https://go.dev/dl/go1.20.5.linux-amd64.tar.gz|100203442|application/x-gzip|invalid-last-modified",
			},
			want: "",
			want1: File{
				Path:          "",
				OriginURL:     "",
				ContentLength: 0,
				ContentType:   "",
				LastModified:  time.Time{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseText(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseText() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseText() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
