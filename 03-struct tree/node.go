package main

import "fmt"

// go语言仅支持封装，不支持继承和多态

type treeNode struct {
	value int
	left, right *treeNode
}

// 显示定义和命名方法接受者
func (node treeNode) print() {
	fmt.Println(node.value)
}

// 使用指针作为方法的接收者
// 只有指针才可以改变结构的内容
// nil指针也可以调用方法
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("set value to nil is invalid")
		return
	}
	node.value = value
}

func createNode(value int) *treeNode {
	// 返回了局部变量的地址
	return &treeNode{value: value}
}

func (node *treeNode) traverse() {
	if node ==  nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	// 不论地址还是结构本身，一律使用 . 来访问成员
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	fmt.Println(root, &root.left.right)
	fmt.Println(treeNode{})
	root.print()
	root.right.left.setValue(999)
	root.right.left.print()

	nodes := []treeNode {
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	var pRoot *treeNode
	fmt.Println(pRoot)
	pRoot.setValue(200)
	pRoot = &root
	fmt.Println(pRoot)
	pRoot.setValue(300)
	pRoot.print()

	root.traverse()


}