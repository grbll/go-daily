package reversenodesinkgroup

import . "github.com/grbll/go-daily/singly-linked-list"

func getInFront(head *ListNode, tail *ListNode) {
	tail.Next, tail.Next.Next, head.Next = tail.Next.Next, head.Next, tail.Next
}

func checkKLength(head *ListNode, k int) bool {
	walker := head
	for i := 0; i < k; i++ {
		if walker.Next == nil {
			return false
		} else {
			walker = walker.Next
		}
	}
	return true
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	head = &ListNode{Next: head}
	var dummy *ListNode = head
	if k != 1 {
		var tail *ListNode = head.Next
		for checkKLength(head, k) {
			for i := 0; i < k-1; i++ {
				getInFront(head, tail)
			}
			head, tail = tail, tail.Next
		}
	}
	return dummy.Next
}
