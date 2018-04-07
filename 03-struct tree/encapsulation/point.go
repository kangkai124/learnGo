package point

import "fmt"

/* 封装
 * 名字一般使用CamelCase
 * 首字母大写：public
 * 首字母小写：private
 */

 /* 包
  * 每个目录一个包
  * main包包含可执行入口
  */ 
type Node struct {
	Value int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("set Value to nil is invalid")
		return
	}
	node.Value = Value
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

func (node *Node) Traverse() {
	if node ==  nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

