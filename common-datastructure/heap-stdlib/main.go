package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func main() {
	h := &IntHeap{5, 3, 8, 1, 7}
	heap.Init(h)

	fmt.Println("Min:", (*h)[0])

	heap.Push(h, 2)
	fmt.Println("After push 2:", *h)

	val := heap.Pop(h)
	fmt.Println("Popped:", val)
	fmt.Println("After pop:", *h)
}
