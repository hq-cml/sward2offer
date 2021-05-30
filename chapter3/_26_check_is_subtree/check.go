package _26_check_is_subtree

import "github.com/hq-cml/sward2offer/common"

//递归！
//我自己的方法，比书上的还精简~
func Check(root1, root2 *common.TreeNode) bool {
    if root2 == nil {
        return true
    }
    if root1 == nil {
        return false
    }

    if root1.I == root2.I {
        return (Check(root1.Left, root2.Left) &&
            Check(root1.Right, root2.Right)) ||
            Check(root1.Left, root2) ||
            Check(root1.Right, root2)
    } else {
        return Check(root1.Left, root2) ||
            Check(root1.Right, root2)
    }
}


