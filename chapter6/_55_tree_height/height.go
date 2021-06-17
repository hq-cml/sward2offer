package _55_tree_height

import (
    "github.com/hq-cml/sward2offer/common"
    "math"
)

//树高度
//难度：2*
func TreeHeight(root *common.TreeNode) int {
    if root == nil {
        return 0
    }

    return 1 + func(tree1, tree2 *common.TreeNode) int {
        height1 := TreeHeight(tree1)
        height2 := TreeHeight(tree2)
        if height1 > height2 {
            return height1
        } else {
            return height2
        }
    }(root.Left, root.Right)
}

//判断是否平衡二叉树
//平衡二叉树：递归的概念，树的任意子树高度差<=1
//普通的先序遍历，效率不高
//难度：2*
func CheckBalance(root *common.TreeNode) bool {
    if root == nil {
        return true
    }

    leftHeight := TreeHeight(root.Left)
    rightHeight := TreeHeight(root.Right)
    if int(math.Abs(float64(leftHeight - rightHeight))) > 1 {
        return false
    }

    return CheckBalance(root.Left) && CheckBalance(root.Right)
}

//更加高效的方法
//利用后续遍历，将统计高度和判断平衡同时执行，并且还能避免重复计算
//难度：4*
func CheckBalace2(root *common.TreeNode) bool {
    depth := 0
    return isBalance(root, &depth)
}

func isBalance(root *common.TreeNode, depth *int) bool {
    if root == nil {
        *depth = 0
        return true
    }

    left := 0
    right := 0
    if !isBalance(root.Left, &left) ||
        !isBalance(root.Right, &right) {
        return false
    }

    //高度差
    diff := int(math.Abs(float64(left - right)))
    if diff <= 1 {
        if left > right {
            *depth = 1 + left
        } else {
            *depth = 1 + right
        }
        return true
    }
    return false
}