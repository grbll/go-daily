package permutationsii

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_next(t *testing.T) {
	type args struct {
		positions []int
		end       int
	}
	tests := []struct {
		name  string
		args  args
		want1 []int
		want2 bool
	}{
		{name: "just one position",
			args:  args{positions: []int{1}, end: 2},
			want1: []int{2},
			want2: true,
		},
		{name: "just one at end",
			args:  args{positions: []int{2}, end: 2},
			want1: []int{2},
			want2: false,
		},
		{name: "two, with one step",
			args:  args{positions: []int{1, 2}, end: 3},
			want1: []int{1, 3},
			want2: true,
		},
		{name: "three, with multi, step:1",
			args:  args{positions: []int{0, 2, 2}, end: 2},
			want1: []int{1, 1, 1},
			want2: true,
		},
		{name: "three, with multi, step:2",
			args:  args{positions: []int{1, 1, 1}, end: 2},
			want1: []int{1, 1, 2},
			want2: true,
		},
		{name: "three, with multi, step:3",
			args:  args{positions: []int{1, 1, 2}, end: 2},
			want1: []int{1, 2, 2},
			want2: true,
		},
		{name: "three, with multi, step:4",
			args:  args{positions: []int{1, 2, 2}, end: 2},
			want1: []int{2, 2, 2},
			want2: true,
		},
		{name: "three, with multi, step:5",
			args:  args{positions: []int{2, 2, 2}, end: 2},
			want1: []int{2, 2, 2},
			want2: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got2 := next(tt.args.positions, tt.args.end)
			if !cmp.Equal(tt.args.positions, tt.want1) {
				t.Errorf("next() = %v, want %v", tt.args.positions, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("next() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {name: "just one",
		// 	args: args{nums: []int{4}},
		// 	want: [][]int{{4}},
		// },
		// {name: "two different",
		// 	args: args{nums: []int{3, 4}},
		// 	want: [][]int{{4, 3}, {3, 4}},
		// },
		// {name: "two equal",
		// 	args: args{nums: []int{3, 3}},
		// 	want: [][]int{{3, 3}},
		// },
		// {name: "three different",
		// 	args: args{nums: []int{3, 4, 7}},
		// 	want: [][]int{{7, 4, 3}, {7, 3, 4}, {4, 7, 3}, {3, 7, 4}, {4, 3, 7}, {3, 4, 7}},
		// },
		{name: "three, two equal",
			args: args{nums: []int{3, 7, 7}},
			want: [][]int{{7, 7, 3}, {7, 3, 7}, {3, 7, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
