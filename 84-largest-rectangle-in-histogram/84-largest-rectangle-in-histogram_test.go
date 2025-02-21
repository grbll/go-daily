package largestrectangleinhistogram

import (
	"testing"
	// "github.com/google/go-cmp/cmp"
)

// func Test_split(t *testing.T) {
// 	type args struct {
// 		ranges       *[][2]int
// 		splitIndices []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want [][2]int
// 	}{
// 		{
// 			name: "two to three",
// 			args: args{ranges: &[][2]int{{0, 4}, {4, 7}}, splitIndices: []int{3, 5}},
// 			want: [][2]int{{0, 3}, {4, 5}, {6, 7}},
// 		},
// 		{
// 			name: "one to one",
// 			args: args{ranges: &[][2]int{{0, 4}}, splitIndices: []int{3}},
// 			want: [][2]int{{0, 3}},
// 		},
// 		{
// 			name: "one to two",
// 			args: args{ranges: &[][2]int{{0, 4}}, splitIndices: []int{2}},
// 			want: [][2]int{{0, 2}, {3, 4}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			split(tt.args.ranges, tt.args.splitIndices)
// 			if !cmp.Equal(*tt.args.ranges, tt.want) {
// 				t.Errorf("split() = %v, want %v\n", *tt.args.ranges, tt.want)
// 			}
// 		})
// 	}
// }

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "reset",
			args: args{heights: []int{4, 2, 0, 3, 2, 5}},
			want: 6,
		},
		{
			name: "sink",
			args: args{heights: []int{2, 1, 2}},
			want: 3,
		},
		{
			name: "gap",
			args: args{heights: []int{2, 2, 1, 0, 2, 2, 2}},
			want: 6,
		},
		{
			name: "lower end",
			args: args{heights: []int{2, 2, 1}},
			want: 4,
		},
		{
			name: "rectangel",
			args: args{heights: []int{2, 2, 2}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
