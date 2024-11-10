package swapnodesinpairs

import . "github.com/grbll/go-daily/singly-linked-list"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		var beginning, one, two, three *ListNode
		beginning = &ListNode{Next: head}
		one = beginning
		for one.Next != nil && one.Next.Next != nil {
			two = one.Next
			three = two.Next
			one.Next, two.Next, three.Next = three, three.Next, two
			one = one.Next.Next
		}
		return beginning.Next
	}
}
