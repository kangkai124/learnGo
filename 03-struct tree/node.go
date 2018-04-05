package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	// 返回了局部变量的地址
	return &treeNode{value: value}
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	fmt.Println(root)
	fmt.Println(treeNode{})

	nodes := []treeNode {
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
