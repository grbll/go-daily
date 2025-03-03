package partitionlist

import . "github.com/grbll/go-daily/singly-linked-list"

func partition(head *ListNode, x int) *ListNode {
	var lowStart, highStart *ListNode = &ListNode{Next: head}, &ListNode{Next: nil}
	var low, high *ListNode = lowStart, highStart
	for low.Next != nil {
		if low.Next.Val < x {
			low = low.Next
		} else {
			high.Next, low.Next = low.Next, low.Next.Next
			high.Next.Next, high = nil, high.Next
		}
	}
	low.Next = highStart.Next
	return lowStart.Next
}
