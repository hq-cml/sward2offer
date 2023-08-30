/*
*

	*给你一棵二叉树的根节点，返回该树的 直径 。

二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

两节点之间路径的 长度 由它们之间边数表示。
*/
package _543_btree_distance

import "github.com/hq-cml/sward2offer/common"

// 思路：本质上是树的遍历 + 距离求值
//
//	距离求值：某个节点为根的距离，实际上就是左子树高 + 右子树高
func diameterOfBinaryTree(root *common.TreeNode) int {
	max := 0
	var dfs func(*common.TreeNode)
	dfs = func(root *common.TreeNode) {
		if root == nil {
			return
		}
		// 左子树高 + 右子树高 == 当前root的直径，如果大于max则记录下来
		if treeHeight(root.Right)+treeHeight(root.Left) > max {
			max = treeHeight(root.Right) + treeHeight(root.Left)
		}
		dfs(root.Left)
		dfs(root.Right)
		return
	}

	dfs(root)
	return max
}

func treeHeight(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(treeHeight(root.Left), treeHeight(root.Right))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
