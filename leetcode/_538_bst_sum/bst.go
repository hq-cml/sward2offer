/*
 * 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
 * 提醒一下，二叉搜索树满足下列约束条件：
 * 节点的左子树仅包含键 小于 节点键的节点。
 * 节点的右子树仅包含键 大于 节点键的节点。
 * 左右子树也必须是二叉搜索树。
 *
 * 	  1          4
 * 	2  3      6    3
 */
package _538_bst_sum

import "github.com/hq-cml/sward2offer/common"

// 思路，根据二叉搜索树的概念
//
// 如果一个节点有右子树，则该节点的新值，应该是节点值 + 其右子树所有节点累计和
// 如果一个节点没有右子树，则该节点新值，应该是向上找，直到某个节点，当前节点属于其左子树，然后新值就是节点值 + 找到的父节点的节点累计值
// 这种思路还是太抽象了，进一步简化
// 利用二叉搜索树的性质，其实其中序遍历就是有序的，那么问题是否可以转化成一个数组的累计值？
// 1, 2, 3, 4 =>  10, 8, 7, 4
// 利用这种思路，其实可以搞一个变种的递归，中序遍历，但是是：右，中，左的形式，自然就得到了累计
// 牛逼
func ConvertBST(root *common.TreeNode) *common.TreeNode {
	var bst func(*common.TreeNode)
	var pre int // 相当于累计值
	bst = func(root *common.TreeNode) {
		if root == nil {
			return
		}
		// 递归右
		bst(root.Right)

		// 当前处理
		root.Val = root.Val + pre // 当前新值，是前面累计值 + 当前值
		pre = root.Val            // 当前新值赋值给pre，用于后续累计

		// 递归左
		bst(root.Left)
	}
	bst(root)
	return root
}
