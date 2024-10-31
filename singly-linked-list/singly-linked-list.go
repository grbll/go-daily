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
	if l != nil {
		sb := strings.Builder{}
		sb.WriteByte('{')
		for l.Next != nil {
			sb.WriteString(strconv.Itoa(l.Val))
			sb.WriteByte(',')
			l = l.Next
		}
		sb.WriteString(strconv.Itoa(l.Val))
		sb.WriteByte(',')
		sb.WriteByte('}')
		return sb.String()
	} else {
		return "{}"
	}
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
