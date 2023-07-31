/*
 * wc的tt面试题：计算一个二叉树路径数字之和
 *
 * 例：下面数的路径组成12和13，所以最终返回25（12+13）
 *    1
 *  2   3
 */
package _calc_all_path_sum

import "github.com/hq-cml/sward2offer/common"

// 本质上还是树的深度遍历
func DfsCalc(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	var total int
	dfsCalc(root, 0, &total)
	return total
}

func dfsCalc(root *common.TreeNode, curr int, total *int) {
	if root == nil {
		return
	}
	curr = curr * 10 + root.Val
	// 是叶子
	if root.Left == nil && root.Right == nil {
		*total = *total + curr
	} else {
		dfsCalc(root.Left, curr, total)
		dfsCalc(root.Right, curr, total)
	}
}