/*
 * 面试题36：二叉搜索树与双向链表
 * 题目：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求
 * 不能创建任何新的结点，只能调整树中结点指针的指向。
 */
package _36_conver_tree_to_dlist

import "github.com/hq-cml/sward2offer/common"

//思路：
//P33中有关于二叉搜索树的定义：根据定义，二叉树的中序遍历，其实就是链表的序列
//关于树的题目，首先先搞就是想到递归的思路
//返回双向链表的第一个节点地址
//沿用了treeNode结构，Left表示pre，Right表示post
//难度等级: 5*
func Convert(root *common.TreeNode) *common.TreeNode {
	var currTail *common.TreeNode
	convert(root, &currTail)

	//此时nowTail指向了队列尾部，需要不断向前移动到头部
	head := currTail
	for head != nil && head.Left != nil {
		head = head.Left
	}
	return head
}

//本质上，这是一个中序遍历递归
//但是，由于指针比较多，所以比较抽象（尤其是root.Right的没有做处理，交给递归过程了）
//第二个参数currTail，即当前已经处理好的链表尾部，通过这个二级指针，将整个链表顺直
//难度：5*
func convert(root *common.TreeNode, currTail **common.TreeNode) {
	if root == nil {
		return
	}

	//递归左子树
	if root.Left != nil {
		convert(root.Left, currTail)
	}

	//当前节点是root，处理左节点善后
	root.Left = *currTail
	if (*currTail) != nil {
		(*currTail).Right = root
	}

	//currTail后移成为当前节点（root.Right没有特殊处理，交给下面的递归过程了）
	*currTail = root

	//递归处理右子树
	if root.Right != nil {
		convert(root.Right, currTail)
	}
}
