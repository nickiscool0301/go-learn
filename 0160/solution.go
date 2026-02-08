package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointerA, pointerB := headA, headB
	for pointerA != pointerB {
		if pointerA != nil {
			pointerA = pointerA.Next
		} else {
			pointerA = headB
		}
		if pointerB != nil {
			pointerB = pointerB.Next
		} else {
			pointerB = headA
		}
	}
	return pointerA
}
