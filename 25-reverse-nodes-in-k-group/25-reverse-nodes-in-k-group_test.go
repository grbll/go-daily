package reversenodesinkgroup

import (
	"reflect"
	"testing"

	. "github.com/grbll/go-daily/singly-linked-list"
)

func Test_getInFront(t *testing.T) {
	tests := []struct {
		name string
		list []int
		head int
		tail int
		want []int
	}{
		{
			name: "Move last node to front in 4-node list",
			list: []int{99, 1, 2, 3},
			head: 0,
			tail: 2,
			want: []int{99, 3, 1, 2},
		},
		{
			name: "Move last node to front in 5-node list",
			list: []int{99, 1, 2, 3, 4},
			head: 0,
			tail: 2,
			want: []int{99, 3, 1, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := BuildListNode(tt.list)
			tail := getNodeAtIndex(head, tt.tail)
			getInFront(head, tail)

			if got := BuildSlice(head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getNodeAtIndex(head *ListNode, index int) *ListNode {
	for i := 0; i < index && head != nil; i++ {
		head = head.Next
	}
	return head
}

func Test_checkKLength(t *testing.T) {
	tests := []struct {
		name string
		list []int
		k    int
		want bool
	}{
		{
			name: "List has exact k nodes",
			list: []int{1, 2, 3, 4, 5},
			k:    5,
			want: true,
		},
		{
			name: "List has more than k nodes",
			list: []int{1, 2, 3, 4, 5, 6},
			k:    4,
			want: true,
		},
		{
			name: "List has fewer than k nodes",
			list: []int{1, 2, 3},
			k:    5,
			want: false,
		},
		{
			name: "Empty list",
			list: []int{},
			k:    1,
			want: false,
		},
		{
			name: "Single-node list, k=1",
			list: []int{1},
			k:    1,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := &ListNode{Next: BuildListNode(tt.list)}
			if got := checkKLength(head, tt.k); got != tt.want {
				t.Errorf("checkKLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		list []int
		k    int
		want []int
	}{
		{
			name: "List with multiple groups of k",
			list: []int{1, 2, 3, 4, 5, 6},
			k:    2,
			want: []int{2, 1, 4, 3, 6, 5},
		},
		{
			name: "List with exact multiple of k",
			list: []int{1, 2, 3, 4, 5, 6},
			k:    3,
			want: []int{3, 2, 1, 6, 5, 4},
		},
		{
			name: "List with remainder after full k groups",
			list: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "List with single group of k",
			list: []int{1, 2, 3},
			k:    3,
			want: []int{3, 2, 1},
		},
		{
			name: "List with k=1 (no reversal)",
			list: []int{1, 2, 3, 4},
			k:    1,
			want: []int{1, 2, 3, 4},
		},
		{
			name: "List shorter than k",
			list: []int{1, 2},
			k:    3,
			want: []int{1, 2},
		},
		{
			name: "Empty list",
			list: []int{},
			k:    2,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := BuildListNode(tt.list)
			got := BuildSlice(reverseKGroup(head, tt.k))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
