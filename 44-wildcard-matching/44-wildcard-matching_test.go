package wildcardmatching

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	// "github.com/google/go-cmp/cmp"
)

func Test_merge(t *testing.T) {
	type args struct {
		ranges    []int
		cutter    []int
		newRanges []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "two identical ranges",
			args: args{ranges: []int{0, 2}, cutter: []int{0, 2}, newRanges: []int{}},
			want: []int{1, 3},
		},
		{
			name: "intersection",
			args: args{ranges: []int{0, 3}, cutter: []int{1, 4}, newRanges: []int{}},
			want: []int{2, 4},
		},
		{
			name: "intersection other way",
			args: args{ranges: []int{1, 4}, cutter: []int{0, 3}, newRanges: []int{}},
			want: []int{2, 4},
		},
		{
			name: "inclusion",
			args: args{ranges: []int{1, 5}, cutter: []int{2, 3}, newRanges: []int{}},
			want: []int{3, 4},
		},
		{
			name: "empty",
			args: args{ranges: []int{1, 5}, cutter: []int{5, 6}, newRanges: []int{}},
			want: []int{},
		},
		{
			name: "inclusion, plus irrelevant",
			args: args{ranges: []int{1, 5}, cutter: []int{2, 4, 5, 6}, newRanges: []int{}},
			want: []int{3, 5},
		},
		{
			name: "complex",
			args: args{ranges: []int{1, 7, 8, 10}, cutter: []int{2, 4, 5, 8}, newRanges: []int{}},
			want: []int{3, 5, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.args.ranges, tt.args.cutter, tt.args.newRanges)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("isMatch(%+v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

// func Test_isMatch(t *testing.T) {
// 	type args struct {
// 		s string
// 		p string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{
// 			name: "mixed case, with wildcard",
// 			args: args{s: "acdcb", p: "a*c?b"},
// 			want: false,
// 		},
// 		{
// 			name: "mixed case",
// 			args: args{s: "adceb", p: "*a*b"},
// 			want: true,
// 		},
// 		{
// 			name: "string empty, pattern star",
// 			args: args{s: "", p: "**"},
// 			want: true,
// 		},
// 		{
// 			name: "simple match",
// 			args: args{s: "aa", p: "aa"},
// 			want: true,
// 		},
// 		{
// 			name: "first simple non match",
// 			args: args{s: "aa", p: "a"},
// 			want: false,
// 		},
// 		{
// 			name: "complete match by *",
// 			args: args{s: "aa", p: "*"},
// 			want: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
// 				t.Errorf("isMatch(%+v) = %v, want %v", tt.args, got, tt.want)
// 			}
// 		})
// 	}
// }
