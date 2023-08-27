package util

import "testing"

func TestIsAlpineVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "OK",
			args: args{
				version: "v3.9",
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				version: "latest-stable",
			},
			want: false,
		},
		{
			name: "invalid",
			args: args{
				version: "v3.9.a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlpineVersion(tt.args.version); got != tt.want {
				t.Errorf("IsAlpineVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
