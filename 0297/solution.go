package main

import (
	"strconv"
	"strings"
)

/* Some personal notes
* A pointer (*int, *[]string): holds a memory address
	* &x: give me addres of x
   * *p: give me value at this address
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) preOrder(node *TreeNode, res *[]string) {
	if node == nil {
		*res = append(*res, "N")
		return
	}
	*res = append(*res, strconv.Itoa(node.Val))
	this.preOrder(node.Left, res)
	this.preOrder(node.Right, res)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []string
	this.preOrder(root, &res)
	return strings.Join(res, "#")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, "#")
	index := 0
	return buildTree(nodes, &index)
}

func buildTree(nodes []string, index *int) *TreeNode {
	if nodes[*index] == "N" {
		*index++
		return nil
	}
	val, _ := strconv.Atoi(nodes[*index])
	*index++
	node := &TreeNode{Val: val}
	node.Left = buildTree(nodes, index)
	node.Right = buildTree(nodes, index)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
