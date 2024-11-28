package wildcardmatching

import (
	"testing"
	// "github.com/google/go-cmp/cmp"
)

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string empty, pattern star",
			args: args{s: "", p: "**"},
			want: true,
		},
		{
			name: "simple match",
			args: args{s: "aa", p: "aa"},
			want: true,
		},
		{
			name: "first simple non match",
			args: args{s: "aa", p: "a"},
			want: false,
		},
		{
			name: "complete match by *",
			args: args{s: "aa", p: "*"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch(%+v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
