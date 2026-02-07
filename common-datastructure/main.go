package main

import "fmt"

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
	fmt.Println("---------------------------------------------\n")

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
