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

func main() {
	kindaSet := NewSet[string]()

	kindaSet.Add("apple")
	kindaSet.Add("banana")
	kindaSet.Add("apple")

	kindaSet.Print()
	fmt.Print("\n")
	fmt.Printf("Set size: %v\n", kindaSet.Len())
	fmt.Printf("Contains apple: %v\n", kindaSet.Contains("apple"))
	fmt.Printf("Contains banana: %v\n", kindaSet.Contains("banana"))

	kindaSet.Remove("banana")
	kindaSet.Print()
	fmt.Print("\n")
	fmt.Printf("Contains banana: %v\n", kindaSet.Contains("banana"))

}
