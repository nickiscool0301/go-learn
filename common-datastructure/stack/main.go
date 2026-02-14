package main

import "fmt"

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("pop from empty stack")
	}
	head := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return head
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		panic("peek from empty stack")
	}
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println("Peek:", s.Peek())
	fmt.Println("Pop:", s.Pop())
	fmt.Println("Len:", s.Len())
}
