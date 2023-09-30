/*
 *  给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
 * 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
 */
package _437_path_sum

import "github.com/hq-cml/sward2offer/common"

// 思路：本质上还是树的遍历
func pathSum(root *common.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var ret int
	help(root, targetSum, &ret)

	return ret + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

// 遍历一个单节点
func help(node *common.TreeNode, target int, ret *int) {
	if node == nil {
		return
	}

	if node.Val == target {
		*ret++
	}

	help(node.Left, target-node.Val, ret)
	help(node.Right, target-node.Val, ret)
}
