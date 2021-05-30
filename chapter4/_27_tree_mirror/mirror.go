package _27_tree_mirror

import "github.com/hq-cml/sward2offer/common"

func Mirror(root *common.TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	Mirror(root.Left)
	Mirror(root.Right)
}
