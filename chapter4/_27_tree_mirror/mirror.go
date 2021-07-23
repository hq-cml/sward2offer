/*
 * 面试题27：二叉树的镜像
 * 题目：请完成一个函数，输入一个二叉树，该函数输出它的镜像。
 */
package _27_tree_mirror

import "github.com/hq-cml/sward2offer/common"

//思路：
//树的题目，首先就应该思考递归策略：左右子节点互换，进行比对。
func Mirror(root *common.TreeNode) {
	//结束条件
	if root == nil {
		return
	}

	//左右互换
	root.Left, root.Right = root.Right, root.Left

	//递归下去
	Mirror(root.Left)
	Mirror(root.Right)
}
