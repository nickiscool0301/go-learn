package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{Val: 0}
	dummy := res
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
		dummy = dummy.Next
	}
	if list1 != nil {
		dummy.Next = list1
	} else {
		dummy.Next = list2
	}
	return res.Next
}
