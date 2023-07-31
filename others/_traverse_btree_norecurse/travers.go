/*
 * zyb面试题: 非递归，遍历二叉树
 */
package _traverse_btree_norecurse

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
)

//非递归，先序遍历二叉树
//这种写法比较非主流，但是符合我个人思维
//此外这种写法有一个问题是仅适用于前序，中序和后序都不太适用
func Pre1(root *common.TreeNode) {
	if root == nil {
		return
	}
	stk := common.NewStack(false)
	stk.Push(root) //先将第一个压入栈
	var p *common.TreeNode
	for stk.Len() > 0 {
		top, _ := stk.Pop()
		p = top.(*common.TreeNode)
		fmt.Print(p.Val, " ")
		if p.Right != nil {
			stk.Push(p.Right)
		}
		if p.Left != nil {
			stk.Push(p.Left)
		}
	}
}

//常规标准写法
//关键在于打印的时机，先序是入栈就打印
func Pre2(root *common.TreeNode) {
	if root == nil {
		return
	}
	stk := common.NewStack(false)
	p := root
	for p != nil || stk.Len() > 0 {
		for p != nil {
			fmt.Print(p.Val, " ")
			stk.Push(p)
			p = p.Left
		}

		top, _ := stk.Pop()
		tmp := top.(*common.TreeNode)
		if tmp.Right != nil {
			p = tmp.Right
		}
	}
}

//中序遍历
//关键在于打印的时机，中序是出栈就打印
func Mid(root *common.TreeNode) {
	if root == nil {
		return
	}
	stk := common.NewStack(false)
	p := root
	for p != nil || stk.Len() > 0 {
		for p != nil {
			stk.Push(p)
			p = p.Left
		}

		top, _ := stk.Pop()
		tmp := top.(*common.TreeNode)
		fmt.Print(tmp.Val, " ")

		if tmp.Right != nil {
			p = tmp.Right
		}
	}
}

//后续遍历
//主要是利用了一个栈顶元素获取但不出栈的特性
//关键在于打印的时机，出栈打印并且还需要记录已处理标志
//难度：5*
func Post(root *common.TreeNode) {
	if root == nil {
		return
	}
	stk := common.NewStack(false)
	p := root
	var vistedNode *common.TreeNode // 用于记录上一个刚刚被遍历处理过的节点
	for p != nil || stk.Len() > 0 {
		for p != nil {
			stk.Push(p)
			p = p.Left
		}
		top, _ := stk.Top() //先读出来，暂时不Pop
		tmp := top.(*common.TreeNode)

		if tmp.Right != nil && tmp.Right != vistedNode { //如果节点存在右孩子，并且右孩子没有被遍历过，则先处理右孩子
			p = tmp.Right
		} else {
			//如果没有右孩子，或者是有右孩子但是右孩子刚刚被遍历过，则意味着轮到了自己，理完毕要出栈
			fmt.Print(tmp.Val, " ")
			stk.Pop()
			vistedNode = tmp //记录刚刚处理过的节点
		}
	}
}
