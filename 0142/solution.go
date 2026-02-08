package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow1, fast := head, head
	for fast != nil && fast.Next != nil {
		slow1 = slow1.Next
		fast = fast.Next.Next
		if slow1 == fast {
			slow2 := head
			for slow1 != slow2 {
				slow1 = slow1.Next
				slow2 = slow2.Next
			}
			return slow1
		}
	}
	return nil
}
