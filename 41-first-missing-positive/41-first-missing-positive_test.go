package firstmissingpositive

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_heapRise(t *testing.T) {
	type args struct {
		nums []int
		tail int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Single rise to top",
			args: args{
				nums: []int{3, 6, 5, 1},
				tail: 4,
			},
			want: []int{1, 3, 5, 6}, // 1 rises to the root
		},
		{
			name: "No rise needed",
			args: args{
				nums: []int{1, 3, 6, 2},
				tail: 3,
			},
			want: []int{1, 3, 6, 2},
		},
		{
			name: "Multiple swaps to root",
			args: args{
				nums: []int{4, 8, 9, 2},
				tail: 4,
			},
			want: []int{2, 4, 9, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			got = append(got, tt.args.nums...)
			heapRise(got, tt.args.tail)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("%v: heapRise(%+v)=%v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}

func Test_smallerChild(t *testing.T) {
	type args struct {
		nums []int
		curr int
		tail int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "Two children, left smaller",
			args: args{
				nums: []int{4, 3, 6, 1},
				curr: 0,
				tail: 4,
			},
			want:  1, // index of 3 (left child)
			want1: true,
		},
		{
			name: "Two children, right smaller",
			args: args{
				nums: []int{4, 7, 6, 2},
				curr: 0,
				tail: 4,
			},
			want:  2, // index of 2 (right child)
			want1: true,
		},
		{
			name: "One child only",
			args: args{
				nums: []int{4, 5, 6, 7, 4},
				curr: 1,
				tail: 4,
			},
			want:  3,
			want1: true,
		},
		{
			name: "No children",
			args: args{
				nums: []int{4},
				curr: 0,
				tail: 1,
			},
			want:  0,
			want1: false, // No children
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := smallerChild(tt.args.nums, tt.args.curr, tt.args.tail)
			if got != tt.want {
				t.Errorf("smallerChild(%+v)  = %v, want %v", tt.args, got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("smallerChild(%+v) got1 = %v, want %v", tt.args, got1, tt.want1)
			}
		})
	}
}

func Test_heapSink(t *testing.T) {
	type args struct {
		nums []int
		tail int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Single sink to child",
			args: args{
				nums: []int{6, 3, 5, 8},
				tail: 4,
			},
			want: []int{3, 6, 5, 8}, // 6 sinks to the second position
		},
		{
			name: "No sinking needed",
			args: args{
				nums: []int{1, 3, 6, 8},
				tail: 4,
			},
			want: []int{1, 3, 6, 8}, // Already a valid heap
		},
		{
			name: "with tail at end",
			args: args{
				nums: []int{6, 3, 5, 8, 4},
				tail: 5,
			},
			want: []int{3, 4, 5, 8, 6},
		},
		{
			name: "with short tail",
			args: args{
				nums: []int{6, 3, 5, 8, 4},
				tail: 4,
			},
			want: []int{3, 6, 5, 8, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			got = append(got, tt.args.nums...)
			heapSink(got, tt.args.tail)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("%v: heapRise(%+v)=%v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "All positive consecutive numbers starting from 1",
		// 	args: args{
		// 		nums: []int{1, 2, 3, 4, 5},
		// 	},
		// 	want: 6, // Next missing positive is 6
		// },
		// {
		// 	name: "All positive non-consecutive numbers",
		// 	args: args{
		// 		nums: []int{3, 4, 7, 1},
		// 	},
		// 	want: 2, // Missing 2
		// },
		// {
		// 	name: "Contains zeros and negatives",
		// 	args: args{
		// 		nums: []int{0, -1, -3, 4, 1},
		// 	},
		// 	want: 2, // Missing 2
		// },
		// {
		// 	name: "Unordered numbers with duplicates",
		// 	args: args{
		// 		nums: []int{3, 4, 3, 1, 1},
		// 	},
		// 	want: 2, // Missing 2
		// },
		// {
		// 	name: "Empty array",
		// 	args: args{
		// 		nums: []int{},
		// 	},
		// 	want: 1, // Missing 1
		// },
		// {
		// 	name: "All negative numbers",
		// 	args: args{
		// 		nums: []int{-1, -2, -3, -4},
		// 	},
		// 	want: 1, // Missing 1
		// },
		// {
		// 	name: "Single positive number 1",
		// 	args: args{
		// 		nums: []int{1},
		// 	},
		// 	want: 2, // Missing 2
		// },
		// {
		// 	name: "Single positive number not 1",
		// 	args: args{
		// 		nums: []int{2},
		// 	},
		// 	want: 1, // Missing 1
		// },
		{
			name: "Large unordered array",
			args: args{
				nums: []int{100, 1, 3, 2, 50, 7, 4, 6, 5},
			},
			want: 8, // Next missing positive is 8
		},
		{
			name: "Already complete range",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 11, // Missing 11
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("%v: firstMissingPositive(%+v) = %v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}
