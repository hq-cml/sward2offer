/*
 * 面试题55（一）：二叉树的深度（高度）
 * 题目：输入一棵二叉树的根结点，求该树的深度。从根结点到叶结点依次经过的
 * 结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。
 *
 * 面试题55（二）：平衡二叉树
 * 题目：输入一棵二叉树的根结点，判断该树是不是平衡二叉树。
 * 平衡二叉树定义：如果某二叉树中任意结点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
 */

package _55_tree_height

import (
	"github.com/hq-cml/sward2offer/common"
	"math"
)

// 题目1：树高度
// 思路：看到树，递归是第一思考顺位
// 难度：2*
func TreeHeight(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(
		float64(TreeHeight(root.Left)), float64(TreeHeight(root.Right))))
}

// 题目2：判断是否平衡二叉树
// 平衡二叉树：递归的概念，树的任意子树高度差<=1
// 思路1：递归遍历各个节点，然后用上面的求深度函数，分别对其左右子节点计算高度。但是显然这样非常低效，因为存在大量的重复
// 思路2：作者的思路很巧妙，利用后序遍历+递归的思路，这样既不存在重复计算，也能够快速得到结果。
// 难度：2*
// 思路1：
func CheckBalance(root *common.TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := TreeHeight(root.Left)
	rightHeight := TreeHeight(root.Right)
	if int(math.Abs(float64(leftHeight-rightHeight))) > 1 {
		return false
	}

	return CheckBalance(root.Left) && CheckBalance(root.Right)
}

// 思路2:
// 更加高效的方法
// 利用后续遍历，将统计高度和判断平衡同时执行，并且还能避免重复计算
// 难度：4*
func CheckBalace2(root *common.TreeNode) bool {
	depth := 0
	return isBalance(root, &depth)
}

// 后序递归
// 关键是这个depth参数，很巧妙
func isBalance(root *common.TreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}

	//后序遍历
	left, right := 0, 0
	if !isBalance(root.Left, &left) ||
		!isBalance(root.Right, &right) {
		return false
	}

	//高度差
	diff := int(math.Abs(float64(left - right)))
	if diff <= 1 {
		// 当前节点是满足的，则有必要计算高度以继续
		if left > right { //depth需要选择大的
			*depth = 1 + left
		} else {
			*depth = 1 + right
		}
		return true
	}
	return false
}
