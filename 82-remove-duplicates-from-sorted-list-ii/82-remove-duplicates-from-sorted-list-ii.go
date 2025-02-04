package removeduplicatesfromsortedlistii

import . "github.com/grbll/go-daily/singly-linked-list"

func deleteDuplicates(head *ListNode) *ListNode {
	var prehead *ListNode = &ListNode{Next: head}
	var origin *ListNode = prehead
	for prehead.Next != nil && prehead.Next.Next != nil {
		if prehead.Next.Val == prehead.Next.Next.Val {
			head = prehead.Next.Next
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prehead.Next = head.Next
		} else {
			prehead = prehead.Next
		}
	}
	return origin.Next
}
