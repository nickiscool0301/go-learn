package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    var check_curr_node func(node *TreeNode, low int, high int) bool
    check_curr_node = func(node *TreeNode, low, high int) bool {
        if node == nil {
            return true
        }
        if node.Val <= low || node.Val >= high {
            return false
        }
        return check_curr_node(node.Left, low, node.Val) && check_curr_node(node.Right, node.Val, high)
    }

    return check_curr_node(root, math.MinInt, math.MaxInt)
