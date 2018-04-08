package main

import (
	"../../02-encapsulation"
	"golang.org/x/tools/container/intsets"

	"fmt"
)

// 使用组合扩充类型
type myTreeNode struct {
	node *point.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(100)
	fmt.Println(s.Has(100))
	fmt.Println(s.Has(1000))
}

func main() {
	var root point.Node

	root = point.Node{Value: 3}
	root.Left = &point.Node{}
	root.Right = &point.Node{5, nil, nil}
	root.Right.Left = new(point.Node)
	root.Left.Right = point.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println()

	myNode := myTreeNode{&root}
	myNode.postOrder()

	fmt.Println()
	testSparse()
}
