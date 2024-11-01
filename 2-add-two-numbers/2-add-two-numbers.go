package addtwonumbers

import (
	. "github.com/grbll/go-daily/singly-linked-list"
)

func Get(l **ListNode) int {
	if *l == nil {
		return 0
	} else {
		val := (*l).Val
		*l = (*l).Next
		return val
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result
	value := 0
	carry := 0
	for l1 != nil || l2 != nil {
		carry = value / 10
		value = Get(&l1) + Get(&l2) + carry
		head.Next = &ListNode{Val: value % 10}
		head = head.Next
	}
	if value/10 != 0 {
		head.Next = &ListNode{Val: 1}
	}

	return result.Next
}
