package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeInfo struct {
	Height int
	Parent *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	heightMap := make(map[int]nodeInfo)

	var dfs func(node *TreeNode, height int, parent *TreeNode)
	dfs = func(node *TreeNode, height int, parent *TreeNode) {
		if node == nil {
			return
		}
		heightMap[node.Val] = nodeInfo{Height: height, Parent: parent}
		dfs(node.Left, height+1, node)
		dfs(node.Right, height+1, node)
	}
	dfs(root, 0, nil)
	return heightMap[x].Height == heightMap[y].Height && heightMap[x].Parent != heightMap[y].Parent
}
