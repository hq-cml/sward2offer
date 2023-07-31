/*
 * 面试题26：树的子结构
 * 题目：输入两棵二叉树A和B，判断B是不是A的子结构。
 */
package _26_check_is_subtree

import "github.com/hq-cml/sward2offer/common"

//思路：
//分析，这是一种递归思想，而且是双函数递归。
//首先判断两棵树的根，如果不相等，则分别递归探测树1的左节点和树2以及树1的右节点和树2
//如果相等，则分为两种探测情况：递归判断左右 or 递归探测树1的左节点和树2以及树1的右节点和树2
//难度：4*
func Check(root1, root2 *common.TreeNode) bool {
	//结束条件：root2结束，说明是子树
	if root2 == nil {
		return true
	}
	//结束条件：root1结束，root2未结束，说明不是子树
	if root1 == nil {
		return false
	}

	// 先进行一遍同根校验，如果恰好符合则直接返回
	if checkWithSameRoot(root1, root2) {
		return true
	}

	// 根不符合，那则逐个递归root1的左右各子节点
	return Check(root1.Left, root2) ||
		Check(root1.Right, root2)
}

// 基于同一颗树根，进行判断子树
func checkWithSameRoot(root1, root2 *common.TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}
	return checkWithSameRoot(root1.Left, root2.Left) &&
		checkWithSameRoot(root1.Right, root2.Right)
}