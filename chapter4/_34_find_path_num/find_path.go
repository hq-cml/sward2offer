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
