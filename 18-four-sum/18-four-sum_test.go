package foursum

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{name: "failing testcase from leetcode",
			nums:   []int{-3, -1, 0, 2, 4, 5},
			target: 2,
			want:   [][]int{{-3, -1, 2, 4}},
		},
		{name: "all two",
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want:   [][]int{{2, 2, 2, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
