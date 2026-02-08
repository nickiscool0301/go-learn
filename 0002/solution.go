package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	carry := 0
	res := &ListNode{Val: 0}
	dummy := res

	for l1 != nil || l2 != nil || carry != 0 {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		currSum := val1 + val2 + carry
		digit := currSum % 10
		carry = currSum / 10

		dummy.Next = &ListNode{Val: digit}
		dummy = dummy.Next
	}
	return res.Next
}
