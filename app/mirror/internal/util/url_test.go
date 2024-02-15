package util

import "testing"

func TestDivideURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		{
			name: "",
			args: args{
				url: "https://dl.google.com/go/go1.20.5.darwin-amd64.tar.gz",
			},
			want:  "https://dl.google.com",
			want1: "/go/go1.20.5.darwin-amd64.tar.gz",
			want2: "",
		},
		{
			name: "",
			args: args{
				url: "https://dl.google.com/go/go1.20.5.darwin-amd64.tar.gz?_=1",
			},
			want:  "https://dl.google.com",
			want1: "/go/go1.20.5.darwin-amd64.tar.gz",
			want2: "_=1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := DivideURL(tt.args.url)
			if got != tt.want {
				t.Errorf("DivideURL() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DivideURL() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("DivideURL() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
