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
	//递归结束条件
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}

	//递归
	if root1.Val == root2.Val {
		return (Check(root1.Left, root2.Left) && Check(root1.Right, root2.Right)) || // 同时判断左右，是&&
			Check(root1.Left, root2) || // ||操作用于分别探测的情况
			Check(root1.Right, root2)
	} else {
		return Check(root1.Left, root2) ||
			Check(root1.Right, root2)
	}
}
