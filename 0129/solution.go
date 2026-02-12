package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type nodeInfo struct {
		Node       *TreeNode
		CurrentVal int
	}

	stack := []nodeInfo{{Node: root, CurrentVal: 0}}
	total := 0

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node, curr := top.Node, top.CurrentVal

		curr = 10*curr + node.Val
		if node.Left == nil && node.Right == nil {
			total += curr
			continue
		}
		if node.Left != nil {
			stack = append(stack, nodeInfo{Node: node.Left, CurrentVal: curr})
		}
		if node.Right != nil {
			stack = append(stack, nodeInfo{Node: node.Right, CurrentVal: curr})
		}
	}
	return total
}
