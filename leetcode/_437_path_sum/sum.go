/*
*
  - 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*/
package _437_path_sum

import "github.com/hq-cml/sward2offer/common"

var (
	ret int
)

func PathSum(root *common.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	help(root, targetSum)
	PathSum(root.Left, targetSum)
	PathSum(root.Right, targetSum)
	return ret
}

func help(node *common.TreeNode, target int) {
	if node == nil {
		return
	}

	if node.Val == target {
		ret++
	}

	help(node.Left, target-node.Val)
	help(node.Right, target-node.Val)

}
