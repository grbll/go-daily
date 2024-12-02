package permutations

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {name: "just one element",
		// 	args: args{nums: []int{1}},
		// 	want: [][]int{{1}},
		// },
		{name: "two elements",
			args: args{nums: []int{1, 2}},
			want: [][]int{{1, 2}, {2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
