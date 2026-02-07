package main

import "fmt"

type KindaSet[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *KindaSet[T] {
	return &KindaSet[T]{
		items: make(map[T]struct{}),
	}
}

func (s *KindaSet[T]) Add(val T) {
	s.items[val] = struct{}{}
}

func (s *KindaSet[T]) Contains(val T) bool {
	_, exists := s.items[val]
	return exists
}

func (s *KindaSet[T]) Remove(val T) {
	delete(s.items, val)
}

func (s *KindaSet[T]) Len() int {
	return len(s.items)
}

func (s *KindaSet[T]) Print() {
	for value, _ := range s.items {
		fmt.Printf("%v\t", value)
	}
}
