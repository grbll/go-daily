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
		{
			name: "all positive numbers - no solution",
			nums: []int{1, 2, 3, 4, 5},
			want: [][]int{},
		},
		{
			name: "all negative numbers - no solution",
			nums: []int{-5, -4, -3, -2, -1},
			want: [][]int{},
		},
		{
			name: "duplicates with zero triplet sum",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "mix of positive, negative and zero",
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			name: "no elements - empty list",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "single element - no triplets",
			nums: []int{0},
			want: [][]int{},
		},
		{
			name: "two elements - no triplets",
			nums: []int{-1, 1},
			want: [][]int{},
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
