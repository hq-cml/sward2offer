package common

import "fmt"

//binary tree
type TreeNode struct {
	I     int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(i int) *TreeNode {
	return &TreeNode{
		I: i,
	}
}

func NewNodeWithChild(i int, left, right *TreeNode) *TreeNode {
	return &TreeNode{
		I:     i,
		Left:  left,
		Right: right,
	}
}

func Pre(n *TreeNode) {
	if n == nil {
		return
	}

	fmt.Print(n.I, " ")
	Pre(n.Left)
	Pre(n.Right)
}

func Mid(n *TreeNode) {
	if n == nil {
		return
	}
	Mid(n.Left)
	fmt.Print(n.I, " ")
	Mid(n.Right)
}

func Post(n *TreeNode) {
	if n == nil {
		return
	}
	Post(n.Left)
	Post(n.Right)
	fmt.Print(n.I, " ")
}
