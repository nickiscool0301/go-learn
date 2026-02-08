package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev, curr := (*ListNode)(nil), slow.Next
	for curr != nil {
		nextt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextt
	}
	slow.Next = nil
	head1, head2 := head, prev
	for head2 != nil {
		nextt := head1.Next
		head1.Next = head2
		head1 = head2
		head2 = nextt
	}
}
