package minimumwindowsubstring

import (
	"testing"
)

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "more complex2",
			args: args{s: "aaabbaaba", t: "abbb"},
			want: "bbaab",
		},
		{
			name: "more complex",
			args: args{s: "ABEBANC", t: "ABC"},
			want: "BANC",
		},
		{
			name: "more complex",
			args: args{s: "ADOBECODEBANC", t: "ABC"},
			want: "BANC",
		},
		{
			name: "one to one, one letter",
			args: args{s: "A", t: "A"},
			want: "A",
		},
		{
			name: "one to one letter",
			args: args{s: "BA", t: "A"},
			want: "A",
		},
		{
			name: "one to one",
			args: args{s: "AB", t: "AB"},
			want: "AB",
		},
		{
			name: "whole string",
			args: args{s: "ACCB", t: "AB"},
			want: "ACCB",
		},
		{
			name: "string with extra",
			args: args{s: "CACCBCCC", t: "AB"},
			want: "ACCB",
		},
		{
			name: "string with extra",
			args: args{s: "CACCACCC", t: "AA"},
			want: "ACCA",
		},
		{
			name: "string with extra",
			args: args{s: "CACCADBCC", t: "BAA"},
			want: "ACCADB",
		},
		{
			name: "string with extra",
			args: args{s: "CACCADBCC", t: "BAA"},
			want: "ACCADB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
