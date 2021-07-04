/*
 * 面试题28：对称的二叉树
 * 题目：请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和
 * 它的镜像一样，那么它是对称的。

这道题主要还是考察递归的理解，左右子节点互换，进行比对。
 */
package _28_symmetry_tree

import (
	"github.com/hq-cml/sward2offer/common"
)

func Symmetry(root *common.TreeNode) bool {
	return symmetry(root, root)
}

func symmetry(root1, root2 *common.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return symmetry(root1.Left, root2.Right) &&
		symmetry(root1.Right, root2.Left)
}
