/**
 * 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
 * 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
 * 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
 *
 * 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
 */
package _337_rob

import (
	"github.com/hq-cml/sward2offer/common"
)

// 说白了，就是一颗二叉树，从根到每个子节点构成一个序列，进而退化成抢劫2的题目
//
//	每一个子序列分别取偷，然后求最大的可能性
//
// 思路：这题是动态规划，但是用递归更加简单，
//
//	此外需要考虑到性能问题，也就是要引入记忆性，用一个map来记录已经计算的结果，避免重复计算
func Rob(root *common.TreeNode) int {
	// 难点在于此，如何设计这个map来暂存计算结果
	// map的意义：key表示node节点的父亲节点
	parentRobbed := map[*common.TreeNode]int{}    // 暂存父节点被偷的结果
	parentNotRobbed := map[*common.TreeNode]int{} // 暂存父节点未被偷的结果
	return helper(root, false, parentRobbed, parentNotRobbed)
}

// robFather，父亲是否被偷过
func helper(node *common.TreeNode, robFather bool, parentRobbed, parentNotRobbed map[*common.TreeNode]int) int {
	if node == nil {
		return 0
	}

	if robFather {
		// 父亲被偷，则当前不能偷了
		if v, ok := parentRobbed[node]; ok {
			return v
		}
		curr := helper(node.Left, false, parentRobbed, parentNotRobbed) +
			helper(node.Right, false, parentRobbed, parentNotRobbed)
		parentRobbed[node] = curr
		return curr
	} else {
		// 父亲未被偷，则当前节点可偷可不偷
		if v, ok := parentNotRobbed[node]; ok {
			return v
		}
		yes := node.Val + helper(node.Left, true, parentRobbed, parentNotRobbed) +
			helper(node.Right, true, parentRobbed, parentNotRobbed)
		no := helper(node.Left, false, parentRobbed, parentNotRobbed) +
			helper(node.Right, false, parentRobbed, parentNotRobbed)
		curr := Max(yes, no)
		parentNotRobbed[node] = curr
		return curr
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
