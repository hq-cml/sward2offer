/*
 * Given a binary tree root, a node  in the tree is named good
 * if in the path from root to X there are nonodes with a value
 * greater than X.
 * Return the number of good nodes in the binary tree
 *
 *         3
 *     1       4
 *   3 nil    1  5
 *
 Input: root =[3,1,4,3,null,1.5]Output: 4
 Explanation: Nodes in blue are good
 Root Node (3)is always a good node
 Node 4 -> (3,4) is the maximum value in the path starting from the root.
 Node 5 ->(3, 4,5)is the maximum value in the path
 Node 3 -> (3,1,3) is the maximum value in the path
 */
package _1448_good_node

import "github.com/hq-cml/sward2offer/common"

// 本质上还是先序遍历
// 维护一个到目前为止路径最大值
func Dfs(root *common.TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *common.TreeNode, preMax int) int {
	if root == nil {
		return 0
	}
	num := 0
	if preMax <= root.Val {
		num = 1
		preMax = root.Val
	}
	return num + dfs(root.Left, preMax) + dfs(root.Right, preMax)
}
