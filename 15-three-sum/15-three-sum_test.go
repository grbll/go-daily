package threesum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "standard case with multiple solutions",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
