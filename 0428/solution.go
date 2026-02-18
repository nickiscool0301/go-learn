package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Val      int
	Children []*Node
}

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (this *Codec) preOrder(node *Node, res *[]string) {
	*res = append(*res, strconv.Itoa(node.Val))
	*res = append(*res, strconv.Itoa(len(node.Children)))
	for _, child := range node.Children {
		this.preOrder(child, res)
	}
}
func (this *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}
	var res []string
	this.preOrder(root, &res)
	return strings.Join(res, "#")
}

func (this *Codec) buildTree(nodes []string, index *int) *Node {
	val, _ := strconv.Atoi(nodes[*index])
	*index++
	childrenCount, _ := strconv.Atoi(nodes[*index])
	*index++

	node := &Node{Val: val, Children: []*Node{}}
	for i := 0; i < childrenCount; i++ {
		child := this.buildTree(nodes, index)
		node.Children = append(node.Children, child)
	}
	return node
}

func (this *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	nodes := strings.Split(data, "#")
	index := 0
	return this.buildTree(nodes, &index)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
