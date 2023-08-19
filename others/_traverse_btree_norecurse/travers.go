/*
 * zyb面试题: 非递归，遍历二叉树
 */
package _traverse_btree_norecurse

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
)

// 常规标准写法
// 关键在于打印的时机，先序是入栈就打印
func PreTree(root *common.TreeNode) {
	if root == nil {
		return
	}
	var stk []*common.TreeNode
	p := root
	for p != nil || len(stk) > 0 {
		for p != nil {
			stk = append(stk, p) // push
			fmt.Println(p.Val)
			p = p.Left
		}
		node := stk[len(stk)-1] // pop
		stk = stk[:len(stk)-1]
		if node.Right != nil {
			p = node.Right
		}
	}
}

// 中序遍历
// 关键在于打印的时机，中序是出栈就打印
func MidTree(root *common.TreeNode) {
	if root == nil {
		return
	}
	var stk []*common.TreeNode
	p := root
	for p != nil || len(stk) > 0 {
		for p != nil {
			stk = append(stk, p) // push
			p = p.Left
		}
		node := stk[len(stk)-1] // pop
		stk = stk[:len(stk)-1]
		fmt.Println(node.Val)
		if node.Right != nil {
			p = node.Right
		}
	}
}

// 后续遍历
// 主要是利用了一个栈顶元素获取但不出栈的特性
// 关键在于打印的时机，出栈打印并且还需要记录已处理标志
// 难度：5*
func PostTree(root *common.TreeNode) {
	if root == nil {
		return
	}
	var stk []*common.TreeNode
	var pre *common.TreeNode
	p := root
	for p != nil || len(stk) > 0 {
		for p != nil {
			stk = append(stk, p) // push
			p = p.Left
		}
		node := stk[len(stk)-1] // top

		if node.Right != nil && node.Right != pre {
			//如果节点存在右孩子，并且右孩子没有被遍历过，则先处理右孩子
			p = node.Right
		} else {
			//如果没有右孩子，或者是有右孩子但是右孩子刚刚被遍历过，则意味着轮到了自己，理完毕要出栈
			stk = stk[:len(stk)-1] // pop
			fmt.Println(node.Val)
			pre = node
		}
	}
}
