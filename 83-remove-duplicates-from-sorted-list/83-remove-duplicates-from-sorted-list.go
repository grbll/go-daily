package removeduplicatesfromsortedlist

import . "github.com/grbll/go-daily/singly-linked-list"

func deleteDuplicates(head *ListNode) *ListNode {
	var origin *ListNode = &ListNode{Next: head}
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return origin.Next
}
