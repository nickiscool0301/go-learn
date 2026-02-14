package main

import "fmt"

type Heap[T any] struct {
	items []T
	less  func(a, b T) bool
}

func NewHeap[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		items: make([]T, 0),
		less:  less,
	}
}

func (h *Heap[T]) Len() int {
	return len(h.items)
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Len() == 0
}

func (h *Heap[T]) Peek() T {
	return h.items[0]
}

func (h *Heap[T]) Push(val T) {
	h.items = append(h.items, val)
	h.siftUp(len(h.items) - 1)
}

func (h *Heap[T]) Pop() T {
	n := len(h.items)
	if n == 0 {
		panic("pop from empty heap")
	}

	root := h.items[0]
	last := h.items[n-1]
	h.items = h.items[:n-1]

	if len(h.items) > 0 {
		h.items[0] = last
		h.siftDown(0)
	}

	return root
}

func (h *Heap[T]) Heapify(arr []T) {
	h.items = arr
	n := len(h.items)

	for i := (n - 1) / 2; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *Heap[T]) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if !h.less(h.items[i], h.items[parent]) {
			break
		}
		h.items[i], h.items[parent] = h.items[parent], h.items[i]
		i = parent
	}
}

func (h *Heap[T]) siftDown(i int) {
	n := len(h.items)

	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.less(h.items[left], h.items[smallest]) {
			smallest = left
		}

		if right < n && h.less(h.items[right], h.items[smallest]) {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.items[i], h.items[smallest] = h.items[smallest], h.items[i]
		i = smallest
	}
}

func (h *Heap[T]) Print() {
	for _, value := range h.items {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

func main() {
	minHeap := NewHeap[int](func(a, b int) bool {
		return a < b
	})

	arr := []int{1, 5, 10, 2, 3, 2, 15, 25, 100}
	minHeap.Heapify(arr)
	minHeap.Print()
	minHeap.Push(-1)
	minHeap.Print()
	minHeap.Pop()
	minHeap.Print()
}
