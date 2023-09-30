/*
 * 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
 * 有效 二叉搜索树定义如下：
 * 节点的左子树只包含 小于 当前节点的数。
 * 节点的右子树只包含 大于 当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 * 例子：输入是按层来体现的
 * 输入：root = [2,1,3] => 输出：true
 * 输入：root = [5,1,4,null,null,3,6] => 输出：false
 * 输入：[5,4,6,null,null,3,7] => false
 * 解释：根节点的值是 5 ，但是右子节点的值是 4 。
 */
package _098_check_bsearch_tree

import (
	"github.com/hq-cml/sward2offer/common"
	"math"
)

// 这题很容易就进入误区，只判断根节点和左右子节点
// 实际上应该是根节点和他的左右子树！
func IsValidBST(root *common.TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && !check(root.Left, math.MinInt, root.Val) {
		return false
	}

	if root.Right != nil && !check(root.Right, root.Val, math.MaxInt) {
		return false
	}

	return IsValidBST(root.Left) && IsValidBST(root.Right)
}

// 判断一棵树的所有子树，均处于一个区间内！
func check(root *common.TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	// 这里注意，根据定义，如果==，也是不合法的二叉搜索树
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return check(root.Left, lower, upper) && check(root.Right, lower, upper)
}
