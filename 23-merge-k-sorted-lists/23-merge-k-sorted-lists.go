package mergeksortedlists

import (
	. "github.com/grbll/go-daily/singly-linked-list"
	"slices"
)

func insert(lists []*ListNode, node *ListNode) []*ListNode {
	if node == nil {
		return lists
	} else {
		up := len(lists)
		down := -1
		var pivot int
		for down+1 < up {
			pivot = down + (up-down)/2
			if lists[pivot].Val <= node.Val {
				down = pivot
			} else {
				up = pivot
			}
		}
		return slices.Insert(lists, down+1, node)
	}
}

func pop(lists *[]*ListNode) *ListNode {
	out := (*lists)[0]
	*lists = insert((*lists)[1:], out.Next)
	return out
}

func mergeKLists(lists []*ListNode) *ListNode {
	beginning := &ListNode{Next: nil}
	head := beginning
	workLists := []*ListNode{}
	for _, node := range lists {
		workLists = insert(workLists, node)
	}
	for len(workLists) != 0 {
		head.Next = pop(&workLists)
		head = head.Next
	}
	return beginning.Next
}
