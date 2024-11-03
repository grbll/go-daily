package medianoftwosortedarrays

import (
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "both arrays empty", args: args{nums1: []int{}, nums2: []int{}}, want: 0.0},
		{name: "first array empty", args: args{nums1: []int{}, nums2: []int{1, 2, 3, 4}}, want: 2.5},  // median of [1, 2, 3, 4]
		{name: "second array empty", args: args{nums1: []int{1, 2, 3, 4}, nums2: []int{}}, want: 2.5}, // median of [1, 2, 3, 4]
		{name: "odd total elements", args: args{nums1: []int{1, 3}, nums2: []int{2, 4, 5}}, want: 3.0},
		{name: "even total elements", args: args{nums1: []int{1, 2}, nums2: []int{3, 4, 5, 6}}, want: 3.5},
		{name: "first array shorter", args: args{nums1: []int{1, 5}, nums2: []int{2, 3, 4, 6}}, want: 3.5},
		{name: "duplicate elements", args: args{nums1: []int{1, 2}, nums2: []int{2, 3, 4, 5}}, want: 2.5},
		{name: "negative and positive numbers", args: args{nums1: []int{-3, -1}, nums2: []int{0, 1, 4, 5}}, want: 0.5},
		{name: "one large element", args: args{nums1: []int{1000}, nums2: []int{1, 2, 3, 4, 5}}, want: 3.5},
		{name: "complex case with larger slices", args: args{nums1: []int{1, 3, 8, 10, 12}, nums2: []int{7, 9, 11, 13, 14, 15}}, want: 10.0},
		{name: "problem in submission", args: args{nums1: []int{1, 2, 3}, nums2: []int{1, 2, 2}}, want: 2.0},
		{name: "problem in submission2", args: args{nums1: []int{1, 4, 5}, nums2: []int{2, 3, 6}}, want: 3.5},
		{name: "problem in submission3", args: args{nums1: []int{4, 5}, nums2: []int{1, 2, 3, 6}}, want: 3.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays(%+v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
func Test_findMedianNumAndSortedArray(t *testing.T) {
	type args struct {
		num  int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "num is less than all elements (odd length)", args: args{num: 1, nums: []int{3, 5, 7}}, want: 4.0},
		{name: "num is less than all elements (even length)", args: args{num: 1, nums: []int{4, 6}}, want: 4.0},
		{name: "num is between middle elements (even length)", args: args{num: 5, nums: []int{3, 6}}, want: 5.0},
		{name: "num is in the middle of the array (odd length)", args: args{num: 5, nums: []int{3, 6, 9}}, want: 5.5},
		{name: "num is greater than all elements (odd length)", args: args{num: 10, nums: []int{3, 5, 7}}, want: 6.0},
		{name: "num is greater than all elements (even length)", args: args{num: 8, nums: []int{4, 6}}, want: 6.0},
		{name: "num is the only element in nums", args: args{num: 5, nums: []int{7}}, want: 6.0},
		{name: "num equals an existing element (odd length)", args: args{num: 6, nums: []int{3, 6, 9}}, want: 6.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianNumAndSortedArray(tt.args.num, tt.args.nums); got != tt.want {
				t.Errorf("findMedianNumAndSortedArray(%+v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
func Test_findMedianSortedArraysPreordered(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "one element in first array", args: args{nums1: []int{1}, nums2: []int{2, 3}}, want: 2.0},
		{name: "odd total elements", args: args{nums1: []int{1, 3}, nums2: []int{2, 4, 5}}, want: 3.0},
		{name: "even total elements", args: args{nums1: []int{1, 2}, nums2: []int{3, 4, 5, 6}}, want: 3.5},
		{name: "first array shorter", args: args{nums1: []int{1, 5}, nums2: []int{2, 3, 4, 6}}, want: 3.5},
		{name: "duplicate elements", args: args{nums1: []int{1, 2}, nums2: []int{2, 3, 4, 5}}, want: 2.5},
		{name: "negative and positive numbers", args: args{nums1: []int{-3, -1}, nums2: []int{0, 1, 4, 5}}, want: 0.5},
		{name: "one large element", args: args{nums1: []int{1000}, nums2: []int{1, 2, 3, 4, 5}}, want: 3.5},
		{name: "complex case with larger slices", args: args{nums1: []int{1, 3, 8, 10, 12}, nums2: []int{7, 9, 11, 13, 14, 15}}, want: 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArraysPreordered(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArraysPreordered(%+v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
