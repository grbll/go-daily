package rotatelist

import . "github.com/grbll/go-daily/singly-linked-list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head != nil {
		// var head = &ListNode{Next: head}
		var start, drag *ListNode = &ListNode{Next: head}, head
		for i := 1; i <= k; i++ {
			if head.Next == nil {
				k %= i
				i = 0
				head = drag
			} else {
				head = head.Next
			}
		}
		if k != 0 {
			for head.Next != nil {
				head = head.Next
				drag = drag.Next
			}
			head.Next = start.Next
			head = drag.Next
			drag.Next = nil
			return head
		}
	}
	return head
}
