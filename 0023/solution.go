package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	curr := dummy
	minHeap := &NodeHeap{}

	for _, node := range lists {
		if node != nil {
			heap.Push(minHeap, node)
		}
	}

	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(*ListNode)
		curr.Next = node
		curr = curr.Next
		if node.Next != nil {
			heap.Push(minHeap, node.Next)
		}
	}

	return dummy.Next
}
