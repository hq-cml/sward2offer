/*
 * 给你两棵二叉树： root1 和 root2 。
 * 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。
 * 你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；
 * 否则，不为 null 的节点将直接作为新二叉树的节点。返回合并后的二叉树。
 * 注意: 合并过程必须从两个树的根节点开始。
 */
package _617_merge_btree

import "github.com/hq-cml/sward2offer/common"

// 思路：本质上，还是递归遍历
func MergeTrees(root1 *common.TreeNode, root2 *common.TreeNode) *common.TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = MergeTrees(root1.Left, root2.Left)
	root1.Right = MergeTrees(root1.Right, root2.Right)
	return root1
}
