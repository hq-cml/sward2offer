/*
 * 面试题28：对称的二叉树
 * 题目：请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和
 * 它的镜像一样，那么它是对称的。
 */
package _28_symmetry_tree

import (
	"github.com/hq-cml/sward2offer/common"
)

//思路1：首先用上一题的方案，求出镜像，然后分别得到先、中、后序序列。和原树的三序列对比。效率比较低
//思路2；将一棵树，一分为二！利用递归，左右子节点互换，进行比对。
//难度：4*
func Symmetry(root *common.TreeNode) bool {
	return symmetry(root, root)
}

func symmetry(root1, root2 *common.TreeNode) bool {
	//结束条件
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil { //有一个非空，必然不是镜像
		return false
	}

	//根不相同，必然不是镜像
	if root1.Val != root2.Val {
		return false
	}

	//递归比较
	//注意顺序：左子树和右子树对比
	return symmetry(root1.Left, root2.Right) &&
		symmetry(root1.Right, root2.Left)
}

//错误！的思想
//很容易就进入的误区！
//func Symmetry2(root *common.TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	if root.Right != nil && root.Left != nil &&
//		root.Left.Val == root.Right.Val {
//		return Symmetry2(root.Left) && Symmetry2(root.Right)
//	} else if root.Right == nil && root.Left == nil {
//		return true
//	} else {
//		return false
//	}
//}
