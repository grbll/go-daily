package nextpermutation

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Standard case - middle permutation",
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "Last permutation - wraps around to first",
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "Partial reordering required",
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
		{
			name: "Larger case - simple next permutation",
			nums: []int{1, 3, 2},
			want: []int{2, 1, 3},
		},
		{
			name: "Already in highest lexicographic order",
			nums: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Repeated elements small",
			nums: []int{1, 5, 1},
			want: []int{5, 1, 1},
		},
		{
			name: "Repeated elements",
			nums: []int{2, 2, 3, 3},
			want: []int{2, 3, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := append([]int(nil), tt.nums...) // make a copy to keep the original for the test case check
			nextPermutation(numsCopy)
			if !reflect.DeepEqual(numsCopy, tt.want) {
				t.Errorf("nextPermutation(%v) = %v, want %v", tt.nums, numsCopy, tt.want)
			}
		})
	}
}
