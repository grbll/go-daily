package removenthnodefromendoflist

import . "github.com/grbll/go-daily/singly-linked-list"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head != nil {
		slow := &ListNode{Val: 0, Next: head}
		beginning := slow

		for i := n; i > 0; i-- {
			if head != nil {
				head = head.Next
			} else {
				return beginning
			}
		}

		for head != nil {
			head = head.Next
			slow = slow.Next
		}

		slow.Next = slow.Next.Next
		return beginning.Next
	} else {
		return nil
	}
}
