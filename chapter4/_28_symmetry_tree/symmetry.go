package _28_symmetry_tree

import (
	"github.com/hq-cml/sward2offer/common"
)

func Symmetry(root *common.TreeNode) bool {
	return symmetry(root, root)
}

func symmetry(root1, root2 *common.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	if root1.I != root2.I {
		return false
	}

	return symmetry(root1.Left, root2.Right) &&
		symmetry(root1.Right, root2.Left)
}
