/*
// 面试题34：二叉树中和为某一值的路径
// 题目：输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所
// 有路径。从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。


这道题，还是在变相得考察二叉树的遍历！
 */

package _34_find_path_num

import (
    "fmt"
    "github.com/hq-cml/sward2offer/common"
)

func FindPath(root *common.TreeNode, num int) {
    arr := []int{}
    findPath(root, num, arr)
}

func findPath(node *common.TreeNode, num int, arr []int) {
    if node == nil {
        return
    }

    arr = append(arr, node.Val)
    if sum(arr) == num &&
        node.Left == nil && node.Right == nil {
        fmt.Println("The Path:", arr)
    }

    findPath(node.Left, num, arr)
    findPath(node.Right, num, arr)
}

func sum(arr []int) int {
    num := 0
    for _, v := range arr {
        num += v
    }
    return num
}
