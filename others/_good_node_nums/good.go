/*
  - Given a binary tree root, a node in the tree is named good if in the path
  - from root to X there are no nodes with a value greater than X.
  - Return the number of good nodes in the binary tree.
  - Expample:
  - 3
  - 1    4
  - 3 nil 1  5
    *
  - Input: root =[3,1,4,3,null,1.5]Output: 4
    Explanation: Nodes in blue are good
    Root Node (3)is always a good node
    Node 4 -> (3,4) is the maximum value in the path starting from the root.
    Node 5 ->(3, 4,5)is the maximum value in the path
    Node 3 -> (3,1,3) is the maximum value in the path.
  - PS: Leatcode: 1448
*/
package _good_node_nums

import (
	"github.com/hq-cml/sward2offer/common"
	"math"
)

func Good(root *common.TreeNode) int {
	return dfs(root, math.MinInt64)
}

// 本质上还是先序遍历
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
