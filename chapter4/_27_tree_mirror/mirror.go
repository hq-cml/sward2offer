/*
 * 面试题27：二叉树的镜像
 * 题目：请完成一个函数，输入一个二叉树，该函数输出它的镜像。

这道题主要还是考察递归的理解，左右子节点互换，进行比对。
 */
package _27_tree_mirror

import "github.com/hq-cml/sward2offer/common"

func Mirror(root *common.TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	Mirror(root.Left)
	Mirror(root.Right)
}
