package singlylinkedlist

import (
	"strconv"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	sb := strings.Builder{}
	sb.WriteByte('[')
	for l != nil && l.Next != nil {
		sb.WriteString(strconv.Itoa(l.Val))
		sb.WriteRune(' ')
		l = l.Next
	}
	if l != nil {
		sb.WriteString(strconv.Itoa(l.Val))
	}
	sb.WriteByte(']')
	return sb.String()
}

func BuildListNode(values []int) *ListNode {
	if len(values) > 0 {
		result := &ListNode{}
		head := result

		for index := 0; index < len(values); index++ {
			head.Next = &ListNode{}
			head = head.Next
			head.Val = values[index]
		}
		return result.Next
	} else {
		return nil
	}
}

func BuildSlice(head *ListNode) []int {
	var out []int = []int{}
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func (l1 *ListNode) Equal(l2 *ListNode) bool {
	for l1.Next != nil {
		if l1.Val != l2.Val || l2.Next == nil {
			return false
		} else {
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	return l2.Next == nil
}
