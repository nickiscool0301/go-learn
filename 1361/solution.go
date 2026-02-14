package main

import "container/list"

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	root := findRoot(n, leftChild, rightChild)
	if root == -1 {
		return false
	}

	seen := make(map[int]struct{})
	queue := list.New()
	queue.PushBack(root)
	seen[root] = struct{}{}

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(int)

		for _, child := range []int{leftChild[node], rightChild[node]} {
			if child == -1 {
				continue
			}
			if _, exists := seen[child]; exists {
				return false
			}
			seen[child] = struct{}{}
			queue.PushBack(child)
		}
	}

	return len(seen) == n
}

func findRoot(n int, leftChild []int, rightChild []int) int {
	isChild := make([]bool, n)

	for _, child := range leftChild {
		if child != -1 {
			isChild[child] = true
		}
	}
	for _, child := range rightChild {
		if child != -1 {
			isChild[child] = true
		}
	}

	for i, child := range isChild {
		if !child {
			return i
		}
	}
	return -1
}
