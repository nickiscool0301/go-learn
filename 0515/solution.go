package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeInfo struct {
	Node  *TreeNode
	Depth int
}

func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return []int{}
	}
	stack := []nodeInfo{nodeInfo{Node: root, Depth: 0}}

	for len(stack) != 0 {
		currNodeInfo := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node, depth := currNodeInfo.Node, currNodeInfo.Depth
		if currNodeInfo.Depth == len(res) {
			res = append(res, node.Val)
		} else {
			res[depth] = max(res[depth], node.Val)
		}
		if node.Left != nil {
			stack = append(stack, nodeInfo{Node: node.Left, Depth: depth + 1})
		}
		if node.Right != nil {
			stack = append(stack, nodeInfo{Node: node.Right, Depth: depth + 1})
		}
	}
	return res
}
