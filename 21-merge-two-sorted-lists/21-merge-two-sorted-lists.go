package mergetwosortedlists

import . "github.com/grbll/go-daily/singly-linked-list"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Val: -101, Next: nil}
	beginning := head
	for list1 != nil {
		for list2 != nil && list2.Val <= list1.Val {
			head.Next = list2
			head, list2 = head.Next, list2.Next
		}
		head.Next = list1
		head, list1 = head.Next, list1.Next
	}
	head.Next = list2
	return beginning.Next
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//     var dummyNode ListNode = ListNode{-101,list1}
//     var head *ListNode = &dummyNode
//     for head.Next!=nil && list2!=nil {
//         if list2.Val <= head.Next.Val {
//             var temp *ListNode = &ListNode{list2.Val,head.Next}
//             head.Next = temp
//             head = head.Next
//             list2 = list2.Next
//         } else {
//             head = head.Next
//         }
//     }
//     if list2!=nil{
//         head.Next = list2
//     }
//     return dummyNode.Next
// }
