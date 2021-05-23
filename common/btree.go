package common

import "fmt"

//binary tree
type Node struct {
	I     int
	Left  *Node
	Right *Node
}

func NewNode(i int) *Node {
	return &Node{
		I: i,
	}
}

func NewNodeWithChild(i int, left, right *Node) *Node {
	return &Node{
		I:     i,
		Left:  left,
		Right: right,
	}
}

func Pre(n *Node) {
	if n == nil {
		return
	}

	fmt.Print(n.I, " ")
	Pre(n.Left)
	Pre(n.Right)
}

func Mid(n *Node) {
	if n == nil {
		return
	}
	Mid(n.Left)
	fmt.Print(n.I, " ")
	Mid(n.Right)
}

func Post(n *Node) {
	if n == nil {
		return
	}
	Post(n.Left)
	Post(n.Right)
	fmt.Print(n.I, " ")
}
