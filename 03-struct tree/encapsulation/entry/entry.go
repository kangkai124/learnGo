package main

import "../../encapsulation"

func main() {
	var root point.Node

	root = point.Node{Value: 3}
	root.Left = &point.Node{}
	root.Right = &point.Node{5, nil, nil}
	root.Right.Left = new(point.Node)
	root.Left.Right = point.CreateNode(2)
	root.Right.Left.SetValue(4 )

	root.Traverse()
}
