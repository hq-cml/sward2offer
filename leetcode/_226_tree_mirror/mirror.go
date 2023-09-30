/*
 * 面试题27：二叉树的镜像
 * 题目：请完成一个函数，输入一个二叉树，该函数输出它的镜像。
 */
package _226_tree_mirror

import "github.com/hq-cml/sward2offer/common"

// 思路：
// 树的题目，首先就应该思考递归策略：
// 左右子节点互换，进行比对。
// 难度：2*
func Mirror(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	// 递归处理右子树，并且处理结果将作为左子树
	left := Mirror(root.Right)
	// 类似
	right := Mirror(root.Left)

	// 左右的处理结果，分别赋值给右左
	root.Left, root.Right = left, right
	return root
}
