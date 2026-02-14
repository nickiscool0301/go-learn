/*
The container/list package implements a doubly linked list,
-> supports efficient insertion and deletion from both ends
-> providing the functionality of a deque with O(1) time complexity for these operations.
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	deque := list.New()

	// Push front
	deque.PushFront(1)
	deque.PushFront(2)
	deque.PushFront(3)

	// Push back
	deque.PushBack(4)
	deque.PushBack(5)
	deque.PushBack(6)

	// Print deque
	for e := deque.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("-----------------------------------")

	// Pop front
	fmt.Println(deque.Front().Value)
	deque.Remove(deque.Front())

	// Pop back
	fmt.Println(deque.Back().Value)
	deque.Remove(deque.Back())

	// Print deque after popping
	for e := deque.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
