package main

import "fmt"

type Node struct {
	No    int
	Name  string
	Left  *Node
	Right *Node
}

// 前序遍历 先输出root节点，在root节点左子树，最后root节点右子树
func order(node *Node) {
	if node != nil {
		fmt.Printf("序号是:%d,姓名是:%s \n", node.No, node.Name)
		order(node.Left)
		order(node.Right)
	}
}

// 中序遍历 先输出root节点左子树，在root节点，最后root节点右子树
func Infixorder(node *Node) {
	if node != nil {
		Infixorder(node.Left)
		fmt.Printf("序号是:%d,姓名是:%s \n", node.No, node.Name)
		Infixorder(node.Right)
	}
}

// 后序遍历 先输出root节点左子树，在root节点右子树，最后root节点
func Postorder(node *Node) {
	if node != nil {
		Infixorder(node.Left)
		Infixorder(node.Right)
		fmt.Printf("序号是:%d,姓名是:%s \n", node.No, node.Name)
	}
}
func main() {
	root := &Node{
		No:   1,
		Name: "甘雨",
	}
	left1 := &Node{
		No:   1,
		Name: "神里",
	}
	right1 := &Node{
		No:   1,
		Name: "影",
	}
	left2 := &Node{
		No:   1,
		Name: "胡桃",
	}
	left3 := &Node{
		No:   1,
		Name: "可莉",
	}
	right2 := &Node{
		No:   1,
		Name: "刻晴",
	}
	root.Left = left1
	left1.Left = left2
	left1.Right = right2
	root.Right = right1
	left2.Left = left3
	order(root)
	Postorder(root)
	Infixorder(root)
}
