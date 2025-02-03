package removeduplicatesfromsortedarrayii

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_rise(t *testing.T) {
	type args struct {
		nums  []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "two steps to heap",
			args: args{nums: []int{1, 3, 1, 3, 3, 2, 3, 2}, index: 2},
			want: []int{1, 3, 2, 3, 3, 2, 3, 1},
		},
		{
			name: "preheaped",
			args: args{nums: []int{2, 3, 1}, index: 0},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var numsCopy []int = append([]int{}, tt.args.nums...)
			rise(tt.args.nums, tt.args.index)
			if !cmp.Equal(tt.args.nums, tt.want) {
				t.Errorf("rise(%v) = %v, want %v", numsCopy, tt.args.nums, tt.want)
			}
		})
	}
}

func Test_sink(t *testing.T) {
	type args struct {
		nums []int
		end  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one swap, with irrelevant end",
			args: args{nums: []int{2, 0, 0, 1, 4, 1, 0, 1, 3, 2}, end: 7},
			want: []int{2, 0, 0, 1, 4, 1, 0, 2, 3, 1},
		},
		{
			name: "one swap",
			args: args{nums: []int{1, 3, 2}, end: 0},
			want: []int{2, 3, 1},
		},
		{
			name: "preheaped",
			args: args{nums: []int{2, 3, 1}, end: 0},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var numsCopy []int = append([]int{}, tt.args.nums...)
			sink(tt.args.nums, tt.args.end)
			if !cmp.Equal(tt.args.nums, tt.want) {
				t.Errorf("rise(%v) = %v, want %v", numsCopy, tt.args.nums, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "very complex",
			args: args{nums: []int{-50, -50, -49, -48, -47, -47, -47, -46, -45, -43, -42, -41, -40, -40, -40, -40, -40, -40, -39, -38, -38, -38, -38, -37, -36, -35, -34, -34, -34, -33, -32, -31, -30, -28, -27, -26, -26, -26, -25, -25, -24, -24, -24, -22, -22, -21, -21, -21, -21, -21, -20, -19, -18, -18, -18, -17, -17, -17, -17, -17, -16, -16, -15, -14, -14, -14, -13, -13, -12, -12, -10, -10, -9, -8, -8, -7, -7, -6, -5, -4, -4, -4, -3, -1, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 9, 10, 10, 10, 11, 11, 12, 12, 13, 13, 13, 14, 14, 14, 15, 16, 17, 17, 18, 20, 21, 22, 22, 22, 23, 23, 25, 26, 28, 29, 29, 29, 30, 31, 31, 32, 33, 34, 34, 34, 36, 36, 37, 37, 38, 38, 38, 39, 40, 40, 40, 41, 42, 42, 43, 43, 44, 44, 45, 45, 45, 46, 47, 47, 47, 47, 48, 49, 49, 49, 50}},
			want: []int{-50, -50, -49, -48, -47, -47, -46, -45, -43, -42, -41, -40, -40, -39, -38, -38, -37, -36, -35, -34, -34, -33, -32, -31, -30, -28, -27, -26, -26, -25, -25, -24, -24, -22, -22, -21, -21, -20, -19, -18, -18, -17, -17, -16, -16, -15, -14, -14, -13, -13, -12, -12, -10, -10, -9, -8, -8, -7, -7, -6, -5, -4, -4, -3, -1, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 16, 17, 17, 18, 20, 21, 22, 22, 23, 23, 25, 26, 28, 29, 29, 30, 31, 31, 32, 33, 34, 34, 36, 36, 37, 37, 38, 38, 39, 40, 40, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 47, 47, 48, 49, 49, 50},
		},
		{
			name: "two removal, one single, not heaped",
			args: args{nums: []int{1, 1, 1, 2, 2, 3}},
			want: []int{1, 1, 2, 2, 3},
		},
		{
			name: "two removal, not heaped",
			args: args{nums: []int{1, 1, 1, 2, 2, 2}},
			want: []int{1, 1, 2, 2},
		},
		{
			name: "one removal, not heaped",
			args: args{nums: []int{1, 1, 1, 2, 2}},
			want: []int{1, 1, 2, 2},
		},
		{
			name: "one removal, preheaped",
			args: args{nums: []int{2, 2, 1, 1, 1}},
			want: []int{1, 1, 2, 2},
		},
		{
			name: "one removal",
			args: args{nums: []int{1, 1, 1}},
			want: []int{1, 1},
		},
		{
			name: "sanity test",
			args: args{nums: []int{1, 1}},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var numsCopy []int = append([]int{}, tt.args.nums...)
			removeDuplicates(tt.args.nums)
			if !cmp.Equal(tt.args.nums[:len(tt.want)], tt.want) {
				t.Errorf("rise(%v) = %v, want %v", numsCopy, tt.args.nums[:len(tt.want)], tt.want)
			}
		})
	}
}
