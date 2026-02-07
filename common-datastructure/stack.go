package main

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
