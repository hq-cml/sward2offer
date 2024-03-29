/*
 * 给你二叉树的根结点 root ，请你将它展开为一个单链表：
 *
 * 展开后的单链表应该同样使用 TreeNode，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
 * 展开后的单链表应该与二叉树 先序遍历 顺序相同。
 *
 *   1     1
 * 2  3 =>   2
 *            3
 */
package _114_btree_to_list

import "github.com/hq-cml/sward2offer/common"

// 思路：本质上是先序遍历
func Transfer(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return root
	}

	left := Transfer(root.Left)
	right := Transfer(root.Right)
	root.Left = nil
	root.Right = left
	pre := root
	p := root.Right
	for p != nil {
		pre = p
		p = p.Right
	}
	pre.Right = right
	return root
}
