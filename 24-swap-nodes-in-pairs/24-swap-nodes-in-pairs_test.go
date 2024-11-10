package swapnodesinpairs

import (
	"reflect"
	"testing"

	. "github.com/grbll/go-daily/singly-linked-list"
)

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "Even number of nodes",
			head: BuildListNode([]int{1, 2, 3, 4}),
			want: BuildListNode([]int{2, 1, 4, 3}),
		},
		{
			name: "Odd number of nodes",
			head: BuildListNode([]int{1, 2, 3}),
			want: BuildListNode([]int{2, 1, 3}),
		},
		{
			name: "Single node",
			head: BuildListNode([]int{1}),
			want: BuildListNode([]int{1}),
		},
		{
			name: "Empty list",
			head: BuildListNode([]int{}),
			want: BuildListNode([]int{}),
		},
		{
			name: "Six nodes",
			head: BuildListNode([]int{1, 2, 3, 4, 5, 6}),
			want: BuildListNode([]int{2, 1, 4, 3, 6, 5}),
		},
		{
			name: "Two nodes",
			head: BuildListNode([]int{1, 2}),
			want: BuildListNode([]int{2, 1}),
		},
		{
			name: "Four nodes, identical values",
			head: BuildListNode([]int{1, 1, 1, 1}),
			want: BuildListNode([]int{1, 1, 1, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.head); !reflect.DeepEqual(BuildSlice(got), BuildSlice(tt.want)) {
				t.Errorf("swapPairs() = %v, want %v", BuildSlice(got), BuildSlice(tt.want))
			}
		})
	}
}
